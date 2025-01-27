# Figma Go Client

![Go Version](https://img.shields.io/badge/go-%3E%3D1.22-blue)
![License](https://img.shields.io/github/license/alridev/figma)
![Build Status](https://img.shields.io/github/actions/workflow/status/alridev/figma/release.yml)

> A Go client for interacting with the Figma API.

## Supported Features

Currently, the library supports the following Figma API functionalities:

- **File Operations**: Retrieve information about Figma files, including their structure, nodes, and properties. [GET file](https://www.figma.com/developers/api#get-files-endpoint), [GET file nodes](https://www.figma.com/developers/api#get-file-nodes-endpoint)

- **Image Export**: Export nodes from Figma files into raster formats such as PNG or JPEG. [GET image](https://www.figma.com/developers/api#get-images-endpoint), [GET image fills](https://www.figma.com/developers/api#get-image-fills-endpoint)

- **Font Find**: Getting a static link to download a font file from static.figma.com.

For more detailed information about the Figma REST API and its capabilities, please refer to the [official documentation](https://www.figma.com/developers/api).

Please note that at this time, the library implements only file operations and image export. Support for other Figma API features may be added in future versions.

## Table of Contents

- [Figma Go Client](#figma-go-client)
  - [Supported Features](#supported-features)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Installation](#installation)
  - [Usage](#usage)
  - [API Documentation](#api-documentation)
    - [Constants](#constants)
    - [Variables](#variables)
    - [Functions](#functions)
    - [Types](#types)
      - [`type FigmaAPI`](#type-figmaapi)
      - [`type FigmaCodeError`](#type-figmacodeerror)
      - [`type FigmaError`](#type-figmaerror)
      - [`type FigmaErrorDescription`](#type-figmaerrordescription)
  - [License](#license)
  - [Contact](#contact)

## Overview

This package provides a convenient interface for Go developers to interact with the Figma API, enabling seamless integration of Figma's capabilities into Go applications.

## Installation

To install the package, run:

```bash
go get -u github.com/alridev/figma
```

## Usage

Here's a basic example of how to use the client:

```go
package main

import (
    "log"

    "github.com/alridev/figma"
    "github.com/alridev/figma/files"
    fm "github.com/alridev/figma/models"
)

func main() {
    // Initialize the Figma API client with your access token
    api := figma.NewFigmaAPI("YOUR_FIGMA_ACCESS_TOKEN")

    // Define the file key and node ID
    fileKey := "your_file_key"
    nodeId := "your_node_id"

    // Retrieve the file nodes
    file, err := files.GetFileNodes(api, fileKey, nodeId)
    if err != nil {
        log.Fatalf("Error retrieving file data: %v", err)
    }

    // Find the specific node
    nodeFind := file.Nodes[nodeId]
    if nodeFind.Document.GetActualInstance() == nil {
        log.Fatalf("Node not found: %v", nodeId)
    }

    // Cast the node to a FrameNode
    mainNode := nodeFind.Document.GetActualInstance().(*fm.FrameNode)

    log.Printf("Selected frame: %s - %s of type %s", mainNode.Name, mainNode.Id, mainNode.Type)

    // Access the children of the main node
    nodes := mainNode.Children

    log.Printf("Number of child nodes: %d", len(nodes))

    // Process child nodes as needed
    for _, child := range nodes {
        log.Printf("Child node: %s - %s of type %s", child.GetName(), child.GetId(), child.GetType())
        // Add your processing logic here
    }
}
```

**Explanation:**

1. **Initialization**: Replace `"YOUR_FIGMA_ACCESS_TOKEN"` with your actual Figma API access token. This token is required to authenticate your requests to the Figma API.

2. **File Key and Node ID**: Set `fileKey` to the key of your Figma file and `nodeId` to the ID of the node you want to retrieve. These identifiers are essential for accessing specific parts of your Figma document.

3. **Retrieve File Nodes**: The `files.GetFileNodes` function fetches the nodes from the specified Figma file. If there's an error during this process, the program logs the error and terminates.

4. **Find and Cast Node**: The retrieved node is checked to ensure it exists. It's then cast to a `FrameNode` to access frame-specific properties.

5. **Access Child Nodes**: The `Children` property of the `FrameNode` contains its child nodes. The program logs the number of child nodes and iterates through them, logging their details. You can add your processing logic within this loop to handle each child node as needed.

**Note**: Ensure you have the necessary permissions and that the file and node IDs are correct. The `figma` package should be properly installed and imported in your project.

This example demonstrates how to interact with the Figma API using the `figma` Go package to retrieve and process nodes within a Figma file.

## API Documentation

### Constants

```go
const (
    // BaseURL is the base URL for the Figma API
    BaseURL = "https://api.figma.com"

    // UserAgent is the default User-Agent for API requests
    UserAgent = "Figma-Go-Client-v0.1.0"

    // Timeout is the timeout duration for API requests
    Timeout = 10 * time.Second
)
```

### Variables

```go
var FigmaErrorDescriptions = map[FigmaCodeError]string{
    BadRequest:          "Parameters are invalid or malformed. Please check the input formats. This error can also occur if the requested resources are too large, resulting in a timeout. Please reduce the number and size of objects requested.",
    Forbidden:           "The request was valid, but the server is refusing action. This can happen if the caller does not have the necessary permissions or when making HTTP requests instead of HTTPS requests.",
    NotFound:            "The requested file or resource was not found.",
    RateLimit:           "API requests may be throttled or rate-limited. Please wait a while before attempting the request again (typically a minute). Rate limiting is calculated on a per-user basis. If using an OAuth token, the rate limit is based on the associated user.",
    InternalServerError: "This most commonly occurs for very large image render requests, which may time out our server and return a 500. Please reduce the number and size of objects requested.",
}
```

### Functions

- `func GetResponseBody(response *http.Response) ([]byte, error)`

- `func ParseJsonFromBytes(body []byte, to interface{}) error`

- `func ParseJsonResponse(response *http.Response, to interface{}) error`

### Types

#### `type FigmaAPI`

```go
type FigmaAPI struct {
    // Token is the access token for the Figma API
    Token string

    // BaseURL is the base URL for the Figma API (default: https://api.figma.com)
    BaseURL string

    // UserAgent is the default User-Agent for API requests (default: Figma-Go-Client)
    UserAgent string

    // AdditionalHeaders are additional headers for API requests
    AdditionalHeaders map[string]string
    AdditionalQueries map[string]string

    // Timeout is the timeout duration for API requests (default: 10 seconds)
    Timeout time.Duration
}
```

**Methods:**

- `func NewFigmaAPI(token string) *FigmaAPI`

- `func (api *FigmaAPI) ActivateRequest(request *http.Request)`

- `func (api *FigmaAPI) AddHeader(request *http.Request, key string, value string)`

- `func (api *FigmaAPI) AddHeaders(request *http.Request, headers map[string]string)`

- `func (api *FigmaAPI) AddQueries(request *http.Request, queries map[string]string)`

- `func (api *FigmaAPI) AddQuery(request *http.Request, key string, value string)`

- `func (api *FigmaAPI) GetClient() *http.Client`

- `func (api *FigmaAPI) GetResponse(request *http.Request) (*http.Response, error)`

- `func (api *FigmaAPI) GetResponseBody(response *http.Response) ([]byte, error)`

- `func (api *FigmaAPI) GetResponseNoActivate(request *http.Request) (*http.Response, error)`

- `func (api *FigmaAPI) ParseJsonResponse(response *http.Response, to interface{}) error`

- `func (api *FigmaAPI) SetAdditionalHeaders(headers map[string]string)`

- `func (api *FigmaAPI) SetAdditionalQueries(queries map[string]string)`

- `func (api *FigmaAPI) SetBaseURL(baseURL string)`

- `func (api *FigmaAPI) SetTimeout(timeout time.Duration)`

- `func (api *FigmaAPI) SetUserAgent(userAgent string)`

#### `type FigmaCodeError`

```go
type FigmaCodeError int
```

**Error Codes:**

```go
const (
    BadRequest          FigmaCodeError = 400
    Forbidden           FigmaCodeError = 403
    NotFound            FigmaCodeError = 404
    RateLimit           FigmaCodeError = 429
    InternalServerError FigmaCodeError = 500
)
```

#### `type FigmaError`

```go
type FigmaError struct {
    Status FigmaCodeError `json:"status"`
    Error  bool           `json:"error"`
}
```

#### `type FigmaErrorDescription`

```go
type FigmaErrorDescription struct {
    Code    FigmaCodeError `json:"code"`
    Message string         `json:"message"`
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

- Project Repository: [https://github.com/alridev/figma](https://github.com/alridev/figma)
