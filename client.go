package dovelet

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
			Type:  feature.VisionFeature(),
			Model: "builtin/latest",
			// use default MaxResults
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
