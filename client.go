package pigeon

import (
	"context"
	"fmt"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
	"google.golang.org/api/option"
)

// VisionClient holds the generic properties for any Google Cloud Vision client.
// This type should become a general type for all vision client types in future, for now it only holds some common data.
type VisionClient struct {
	// The context for each client
	context context.Context
	// The credentials for each client
	credentials option.ClientOption
	// Language hints (see https://pkg.go.dev/cloud.google.com/go/vision/v2/apiv1/visionpb#ImageContext.LanguageHints)
	languageHints []string
	// Keep language hints for all requests, once set?
	keepLanguageHints bool
}

// New returns a pointer to a new Client object.
func NewClient(ctx context.Context, credentials option.ClientOption) VisionClient {
	return VisionClient{
		context:     ctx,
		credentials: credentials,
	}
}

// Set language hints for better results (see https://cloud.google.com/vision/docs/languages, the "languageHints code" column).
// Set keep to true to keep the languages for all successive requests.
func (c *VisionClient) SetLanguageHints(languages []string, keep bool) {
	c.languageHints = languages
	c.keepLanguageHints = keep
}

// MakeBatchAnnotateImageRequest performs a batch image annotation request.
func (vc *VisionClient) RequestImageAnnotation(imagePaths []string, feature DetectionFeature) (*visionpb.BatchAnnotateImagesResponse, error) {
	c, err := vision.NewImageAnnotatorClient(vc.context, vc.credentials)
	if err != nil {
		return nil, fmt.Errorf("unable to create an image annotator client: %v", err)
	}
	defer c.Close()

	visionImages := make([]*visionpb.Image, len(imagePaths))

	for i, p := range imagePaths {
		img, err := os.Open(p)
		if err != nil {
			return nil, fmt.Errorf("failed to open %s: %v", p, err)
		}
		defer img.Close()
		// TODO: move to v2 library functions (see https://pkg.go.dev/cloud.google.com/go/vision/v2)
		vimg, err := vision.NewImageFromReader(img)
		if err != nil {
			return nil, fmt.Errorf("failed to create a vision image from %s: %v", p, err)
		}
		visionImages[i] = vimg
	}

	imageRequests := make([]*visionpb.AnnotateImageRequest, len(visionImages))

	for i, vi := range visionImages {
		imageRequests[i] = &visionpb.AnnotateImageRequest{
			Image:    vi,
			Features: make([]*visionpb.Feature, 1),
			ImageContext: &visionpb.ImageContext{
				LanguageHints: vc.languageHints,
			},
		}
		imageRequests[i].Features[0] = &visionpb.Feature{
			Type: feature.VisionFeature(),
			// use default MaxResults and Model
		}
	}

	batchRequest := &visionpb.BatchAnnotateImagesRequest{
		Requests: imageRequests,
		// do not specify Parent and Labels
	}

	if !vc.keepLanguageHints {
		vc.languageHints = []string{}
	}

	return c.BatchAnnotateImages(vc.context, batchRequest)
}

// // ImagesService returns a pointer to a vision's ImagesService object.
// func (c Client) ImagesService() *vision.ImagesService {
// 	return c.service.Images
// }

// // NewBatchAnnotateImageRequest returns a pointer to a new vision's BatchAnnotateImagesRequest.
// func (c Client) NewBatchAnnotateImageRequest(list []string, features ...*vision.Feature) (*vision.BatchAnnotateImagesRequest, error) {
// 	batch := &vision.BatchAnnotateImagesRequest{}
// 	batch.Requests = []*vision.AnnotateImageRequest{}
// 	for _, v := range list {
// 		req, err := c.NewAnnotateImageRequest(v, features...)
// 		if err != nil {
// 			return nil, err
// 		}
// 		batch.Requests = append(batch.Requests, req)
// 	}
// 	return batch, nil
// }

// // NewAnnotateImageRequest returns a pointer to a new vision's AnnotateImagesRequest.
// func (c Client) NewAnnotateImageRequest(v interface{}, features ...*vision.Feature) (*vision.AnnotateImageRequest, error) {
// 	switch v := v.(type) {
// 	case []byte:
// 		// base64
// 		return NewAnnotateImageContentRequest(v, features...)
// 	case string:
// 		u, err := url.Parse(v)
// 		if err != nil {
// 			return nil, err
// 		}
// 		switch u.Scheme {
// 		case "gs":
// 			// GcsImageUri: Google Cloud Storage image URI. It must be in the
// 			// following form:
// 			// "gs://bucket_name/object_name". For more
// 			return NewAnnotateImageSourceRequest(u.String(), features...)
// 		case "http", "https":
// 			httpClient := c.config.HTTPClient
// 			if httpClient == nil {
// 				httpClient = http.DefaultClient
// 			}
// 			resp, err := httpClient.Get(u.String())
// 			if err != nil {
// 				return nil, err
// 			}
// 			defer resp.Body.Close()
// 			if resp.StatusCode >= http.StatusBadRequest {
// 				return nil, http.ErrMissingFile
// 			}
// 			body, err := io.ReadAll(resp.Body)
// 			if err != nil {
// 				return nil, err
// 			}
// 			return c.NewAnnotateImageRequest(body, features...)
// 		}
// 		// filepath
// 		b, err := os.ReadFile(v)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return c.NewAnnotateImageRequest(b, features...)
// 	}
// 	return &vision.AnnotateImageRequest{}, nil
// }

// // NewAnnotateImageContentRequest returns a pointer to a new vision's AnnotateImagesRequest.
// func NewAnnotateImageContentRequest(body []byte, features ...*vision.Feature) (*vision.AnnotateImageRequest, error) {
// 	req := &vision.AnnotateImageRequest{
// 		Image:    NewAnnotateImageContent(body),
// 		Features: features,
// 	}
// 	return req, nil
// }

// // NewAnnotateImageSourceRequest returns a pointer to a new vision's AnnotateImagesRequest.
// func NewAnnotateImageSourceRequest(source string, features ...*vision.Feature) (*vision.AnnotateImageRequest, error) {
// 	req := &vision.AnnotateImageRequest{
// 		Image:    NewAnnotateImageSource(source),
// 		Features: features,
// 	}
// 	return req, nil
// }

// // NewAnnotateImageContent returns a pointer to a new vision's Image.
// // It's contained image content, represented as a stream of bytes.
// func NewAnnotateImageContent(body []byte) *vision.Image {
// 	return &vision.Image{
// 		// Content: Image content, represented as a stream of bytes.
// 		Content: base64.StdEncoding.EncodeToString(body),
// 	}
// }

// // NewAnnotateImageSource returns a pointer to a new vision's Image.
// // It's contained external image source (i.e. Google Cloud Storage image
// // location).
// func NewAnnotateImageSource(source string) *vision.Image {
// 	return &vision.Image{
// 		Source: &vision.ImageSource{
// 			GcsImageUri: source,
// 		},
// 	}
// }
