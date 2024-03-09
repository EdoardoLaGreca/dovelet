# Pigeon

Pigeon is both a Go wrapper module for the Google Cloud Vision API and a tool to use such wrapper.

[![Go Report Card](https://goreportcard.com/badge/github.com/EdoardoLaGreca/pigeon)](https://goreportcard.com/report/github.com/EdoardoLaGreca/pigeon)

The original Pigeon tool and library used deprecated Google Cloud Vision APIs. One reason could be that, at the time of writing, its last commit was on Oct 14, 2020, which means that the APIs may have been deprecated somewhere after the last commit. This fork attempts to stay up to date and to pursue a similar purpose. Apart from the usage of deprecated APIs, I think that the original Pigeon codebase was unnecessarily complex.

## Requirements

You need to export the path of your Service Account private key (it should be a JSON file) as `GOOGLE_APPLICATION_CREDENTIALS` (replace `/path/to/service_account.json`).

```sh
GOOGLE_APPLICATION_CREDENTIALS=/path/to/service_account.json
export GOOGLE_APPLICATION_CREDENTIALS
```

To generate the credentials file refer to ["Authentication with service accounts"](https://cloud.google.com/vision/docs/setup#sa).

## Tool installation

`pigeon` provides the command-line tools.

```sh
go install github.com/EdoardoLaGreca/pigeon/cmd/pigeon@latest
```

Make sure that `pigeon` was installed correctly:

```sh
pigeon -h
```

## Usage

### `pigeon` tool

Use the `pigeon` tool to make requests. The syntax is as follows.

```sh
pigeon -feature [ -lang l1,l2,... ] files ...
```

Where:
- `-feature` is the feature to use (from the available feature set, use `pigeon -h` to get the complete list)
- `l1,l2,...` are language hints for better accuracy (see "`languageHints` code" in [Supported languages](https://cloud.google.com/vision/docs/languages#supported-langs))
- `files ...` is one or more files to use in the request

<!-- TODO: make new gif -->
<!-- ![pigeon-cmd](https://raw.githubusercontent.com/kaneshin/pigeon/main/assets/pigeon-cmd.gif) -->

### `pigeon` package

The pigeon package contains types and functions for simple Google Cloud Vision queries.

Refer to [`cmd/pigeon/main.go`](cmd/pigeon/main.go) for an example.

## Example

input:

![pigeon](https://raw.githubusercontent.com/EdoardoLaGreca/pigeon/main/assets/pigeon.png)

```sh
pigeon -label assets/pigeon.png
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

![lenna](https://raw.githubusercontent.com/EdoardoLaGreca/pigeon/main/assets/lenna.jpg)

```sh
pigeon -safe-search assets/lenna.jpg
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
