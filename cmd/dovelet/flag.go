package main

import (
	"flag"
	"strings"

	"github.com/EdoardoLaGreca/dovelet"
	"github.com/kaneshin/dovelet"
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
	languages       []string
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
	languages := f.String("lang", "", "Specify language hints for text detection (only works for -text and -doc). For more than one hint, separate them using commas `,`, for example \"en,it\".")
	f.Usage = func() {
		f.PrintDefaults()
	}
	f.Parse(args)
	clargs := &CLArgs{
		face:            *faceDetection,
		landmark:        *landmarkDetection,
		logo:            *logoDetection,
		label:           *labelDetection,
		text:            *textDetection,
		docText:         *docTextDetection,
		safeSearch:      *safeSearchDetection,
		imageProperties: *imageProperties,
		languages:       strings.Split(*languages, ","),
		flags:           f,
	}

	if clargs.Feature() != dovelet.TextDetection && clargs.Feature() != dovelet.DocumentTextDetection {
		clargs.languages = []string{}
	}

	return clargs
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
func (d CLArgs) Feature() dovelet.DetectionFeature {
	switch {
	case d.face:
		return dovelet.FaceDetection
	case d.landmark:
		return dovelet.LandmarkDetection
	case d.logo:
		return dovelet.LogoDetection
	case d.label:
		return dovelet.LabelDetection
	case d.text:
		return dovelet.TextDetection
	case d.docText:
		return dovelet.DocumentTextDetection
	case d.safeSearch:
		return dovelet.SafeSearchDetection
	case d.imageProperties:
		return dovelet.ImageProperties
	}
	return dovelet.TypeUnspecified
}

func (d CLArgs) Language() []string {
	return d.languages
}
