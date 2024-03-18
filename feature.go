package dovelet

import (
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
)

type DetectionFeature int

const (
	// Unspecified or unknown feature type.
	TypeUnspecified DetectionFeature = iota
	// Face detection.
	FaceDetection
	// Landmark detection.
	LandmarkDetection
	// Logo detection.
	LogoDetection
	// Label detection.
	LabelDetection
	// OCR with big text.
	TextDetection
	// OCR on document or small text.
	DocumentTextDetection
	// Detect potential sensitive content.
	SafeSearchDetection
	// Compute image properties.
	ImageProperties
)

// DetectionType returns a value of detection type.
func (d DetectionFeature) String() string {
	switch d {
	case TypeUnspecified:
		return "TYPE_UNSPECIFIED"
	case FaceDetection:
		return "FACE_DETECTION"
	case LandmarkDetection:
		return "LANDMARK_DETECTION"
	case LogoDetection:
		return "LOGO_DETECTION"
	case LabelDetection:
		return "LABEL_DETECTION"
	case TextDetection:
		return "TEXT_DETECTION"
	case DocumentTextDetection:
		return "DOCUMENT_TEXT_DETECTION"
	case SafeSearchDetection:
		return "SAFE_SEARCH_DETECTION"
	case ImageProperties:
		return "IMAGE_PROPERTIES"
	}
	return ""
}

func (d DetectionFeature) VisionFeature() visionpb.Feature_Type {
	switch d {
	case TypeUnspecified:
		return visionpb.Feature_TYPE_UNSPECIFIED
	case FaceDetection:
		return visionpb.Feature_FACE_DETECTION
	case LandmarkDetection:
		return visionpb.Feature_LANDMARK_DETECTION
	case LogoDetection:
		return visionpb.Feature_LOGO_DETECTION
	case LabelDetection:
		return visionpb.Feature_LABEL_DETECTION
	case TextDetection:
		return visionpb.Feature_TEXT_DETECTION
	case DocumentTextDetection:
		return visionpb.Feature_DOCUMENT_TEXT_DETECTION
	case SafeSearchDetection:
		return visionpb.Feature_SAFE_SEARCH_DETECTION
	case ImageProperties:
		return visionpb.Feature_IMAGE_PROPERTIES
	}
	return visionpb.Feature_TYPE_UNSPECIFIED
}
