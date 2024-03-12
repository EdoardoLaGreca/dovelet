# Dovelet

Dovelet is both a Go wrapper module for the Google Cloud Vision API and a tool to use such wrapper.

[![build](https://github.com/EdoardoLaGreca/dovelet/actions/workflows/go.yml/badge.svg)](https://github.com/EdoardoLaGreca/dovelet/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/EdoardoLaGreca/dovelet)](https://goreportcard.com/report/github.com/EdoardoLaGreca/dovelet)

The original Pigeon tool and library used deprecated Google Cloud Vision APIs. One reason could be that, at the time of writing, its last commit was on Oct 14, 2020, which means that the APIs may have been deprecated somewhere after the last commit. This fork attempts to stay up to date and to pursue a similar purpose. Apart from the usage of deprecated APIs, I think that the original Pigeon codebase was unnecessarily complex.

## Requirements

You need to export the path of your Service Account private key (it should be a JSON file) as `GOOGLE_APPLICATION_CREDENTIALS` (replace `/path/to/service_account.json`).

```sh
GOOGLE_APPLICATION_CREDENTIALS=/path/to/service_account.json
export GOOGLE_APPLICATION_CREDENTIALS
```

To generate the credentials file refer to ["Authentication with service accounts"](https://cloud.google.com/vision/docs/setup#sa).

## Tool installation

`dovelet` provides the command-line tools.

```sh
go install github.com/EdoardoLaGreca/dovelet/cmd/dovelet@latest
```

Make sure that `dovelet` was installed correctly:

```sh
dovelet -h
```

## Usage

### `dovelet` tool

Use the `dovelet` tool to make requests. The syntax is as follows.

```sh
dovelet -feature [ -lang l1,l2,... ] files ...
```

Where:
- `-feature` is the feature to use (from the available feature set, use `dovelet -h` to get the complete list)
- `l1,l2,...` are language hints for better accuracy (see "`languageHints` code" in [Supported languages](https://cloud.google.com/vision/docs/languages#supported-langs))
- `files ...` is one or more files to use in the request

<!-- TODO: make new gif -->
<!-- ![pigeon-cmd](https://raw.githubusercontent.com/kaneshin/pigeon/main/assets/pigeon-cmd.gif) -->

### `dovelet` package

The dovelet package contains types and functions for simple Google Cloud Vision queries.

Refer to [`cmd/dovelet/main.go`](cmd/dovelet/main.go) for an example.

## Example

input:

![dovelet](https://raw.githubusercontent.com/EdoardoLaGreca/dovelet/main/assets/dovelet.png)

```sh
dovelet -label assets/dovelet.png
```

output:

```json
[
  {
    "labelAnnotations": [
      {
        "description": "bird",
        "mid": "/m/015p6",
        "score": 0.825656
      },
      {
        "description": "anatidae",
        "mid": "/m/01c_0l",
        "score": 0.58264238
      }
    ]
  }
]
```


### Lenna

input:

![lenna](https://raw.githubusercontent.com/EdoardoLaGreca/dovelet/main/assets/lenna.jpg)

```sh
dovelet -safe-search assets/lenna.jpg
```

output:

```json
[
  {
    "safeSearchAnnotation": {
      "adult": "POSSIBLE",
      "medical": "UNLIKELY",
      "spoof": "VERY_UNLIKELY",
      "violence": "VERY_UNLIKELY"
    }
  }
]
```

## License

[MIT](LICENSE)

## Credits

- Author of the original software: [Shintaro Kaneko](https://github.com/kaneshin) ([repo](https://github.com/kaneshin/pigeon))
