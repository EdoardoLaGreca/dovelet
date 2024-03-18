# Contributing

## Adding a Vision feature

To add a Google Cloud Vision feature support to Dovelet (both the utility and the library), you need to change two files: `feature.go` and `cmd/dovelet/flag.go`.

In `feature.go`, you need to:

1. Add a constant to the `DetectionFeature` enumeration. The new constant must be named after the feature to add.
2. Add a `case` in `DetectionFeature`'s `VisionFeature` function. That new case translates the new constant to its respective Vision-specific value.

In `cmd/dovelet/flag.go`, you need to:

1. Add a new boolean field to the `CLArgs` structure.
2. Add a new command line argument in the `ParseArgs` function. To do so:
	1. Create a new variable named after the new feature. Assign the value returned by `f.Bool` to that variable, called with proper argument values.
	2. Assign the variable's value to the new field in the `CLArgs` instance. The value needs to be de-referenced.
3. Add a `case` in `CLArgs`'s `Feature` function. That case maps the new command line argument with the previously added feature (from `feature.go`).
