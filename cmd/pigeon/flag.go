package main

import (
	"flag"

	"github.com/EdoardoLaGreca/pigeon"
)

// Detections type
type CLArgs struct {
	face            bool
	landmark        bool
	logo            bool
	label           bool
	text            bool
	docText         bool
	safeSearch      bool
	imageProperties bool
	language        string
	flags           *flag.FlagSet
}

// TODO: rephrase flag descriptions
// ParseArgs parses the command-line flags from arguments and returns
// a new pointer of a Detections object..
func ParseArgs(args []string) *CLArgs {
	f := flag.NewFlagSet("Detections", flag.ExitOnError)
	faceDetection := f.Bool("face", false, "This flag specifies the face detection of the feature")
	landmarkDetection := f.Bool("landmark", false, "This flag specifies the landmark detection of the feature")
	logoDetection := f.Bool("logo", false, "This flag specifies the logo detection of the feature")
	labelDetection := f.Bool("label", false, "This flag specifies the label detection of the feature")
	textDetection := f.Bool("text", false, "This flag specifies the text detection (OCR) of the feature")
	docTextDetection := f.Bool("doc", false, "This flag specifies the document text detection (OCR) of the feature")
	safeSearchDetection := f.Bool("safe-search", false, "This flag specifies the safe-search of the feature")
	imageProperties := f.Bool("image-properties", false, "This flag specifies the image safe-search properties of the feature")
	language := f.String("lang", "", "Specify a language for text detection (only works for -text and -doc).")
	f.Usage = func() {
		f.PrintDefaults()
	}
	f.Parse(args)
	return &CLArgs{
		face:            *faceDetection,
		landmark:        *landmarkDetection,
		logo:            *logoDetection,
		label:           *labelDetection,
		text:            *textDetection,
		docText:         *docTextDetection,
		safeSearch:      *safeSearchDetection,
		imageProperties: *imageProperties,
		language:        *language,
		flags:           f,
	}
}

// Args returns the non-flag command-line arguments.
func (d CLArgs) Args() []string {
	return d.flags.Args()
}

// Usage prints options of the Detection object.
func (d CLArgs) Usage() {
	d.flags.Usage()
}

// Feature returns the feature specified as a flag.
func (d CLArgs) Feature() pigeon.DetectionFeature {
	switch {
	case d.face:
		return pigeon.FaceDetection
	case d.landmark:
		return pigeon.LandmarkDetection
	case d.logo:
		return pigeon.LogoDetection
	case d.label:
		return pigeon.LabelDetection
	case d.text:
		return pigeon.TextDetection
	case d.docText:
		return pigeon.DocumentTextDetection
	case d.safeSearch:
		return pigeon.SafeSearchDetection
	case d.imageProperties:
		return pigeon.ImageProperties
	}
	return pigeon.TypeUnspecified
}

func (d CLArgs) Language() string {
	return d.language
}
