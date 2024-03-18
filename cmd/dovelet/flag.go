package main

import (
	"flag"
	"strings"

	"github.com/EdoardoLaGreca/dovelet"
)

// A CLArgs holds values command line argument values.
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

// ParseArgs parses the command line flags and arguments.
func ParseArgs(args []string) *CLArgs {
	f := flag.NewFlagSet("Detections", flag.ExitOnError)
	faceDetection := f.Bool("face", false, "Run face detection.")
	landmarkDetection := f.Bool("landmark", false, "Run landmark detection.")
	logoDetection := f.Bool("logo", false, "Run logo detection.")
	labelDetection := f.Bool("label", false, "Run label detection.")
	textDetection := f.Bool("text", false, "Run big text detection (OCR).")
	docTextDetection := f.Bool("doc", false, "Run small/document text detection (OCR).")
	safeSearchDetection := f.Bool("safe-search", false, "Run safe-search detection.")
	imageProperties := f.Bool("image-properties", false, "Compute image properties.")
	languages := f.String("lang", "", "Specify language hints for text detection (only works for -text and -doc). For more than one hint, separate them using commas `,` (e.g. \"en,it\").")
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

// Language returns the language(s) specified as command line arguments for
// text detection.
func (d CLArgs) Language() []string {
	return d.languages
}
