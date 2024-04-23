# go-redcap

> [!WARNING]  
> This is a Go client for the REDCap API. It is a work in progress and is not yet feature complete as of 04/23/2024.

## Overview

It is based on the [REDCap API documentation](https://redcap.vanderbilt.edu/api/help/) and aims to provide a simple and easy-to-use interface for interacting with REDCap projects. Use this client understanding that some functionality may be limited or in development.

## Example Usage

```go
package main

import (
    "fmt"
    "log"

    redcap "github.com/tkruer/go-redcap/pkg"
)

func main() {
    // Create a new client
    client := redcap.RedCapClient{
		URL:            "https://redcap.example.com/api/",
		Token:          "YOUR_API_TOKEN",
		ResponseFormat: "json",
	}
    // Export events from a project
    client.ExportEvents()     
}
```

## Installation

To install the package, run:

```bash
go get github.com/tkruer/go-redcap
```

## Documentation

For more information, please refer to the [GoDoc documentation](https://pkg.go.dev/github.com/tkruer/go-redcap). You can also find additional documentation and examples in the [examples](https://github.com/tkruer/go-redcap/tree/main/examples) directory. Common use cases for this library include exporting data, importing data, and managing records in a REDCap project from either a CLI or a API wrapper.

This library is really a layer of abstraction on top of the REDCap API. It is designed to make it easier to interact with REDCap projects by providing a simple and easy-to-use interface for common tasks. It is not intended to be a complete implementation of the REDCap API, but rather a starting point for building more complex applications.

## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request if you find a bug or would like to suggest an improvement. Before submitting a pull request, please make sure to run `go fmt` and `go test` to ensure that your code is properly formatted and passes all tests.

## Features

Currently, these API calls are available:

Certainly! Below are the markdown tables for the documentation, with a single column for the feature and a checkmark (✅) indicating the capability.

### Export

| Feature                         | Available |
|---------------------------------|:---------:|
| Arms                            |     ✅     |
| Data Access Groups              |     ✅     |
| Events                          |     ✅     |
| Field names                     |     ✅     |
| Instruments                     |     ✅     |
| Instrument-event mapping        |     ✅     |
| File                            |     ✅     |
| Logging                         |     ✅     |
| Metadata                        |     ✅     |
| Project Info                    |     ✅     |
| PDF of instruments              |     ✅     |
| Records                         |     ✅     |
| Repeating instruments and events|     ✅     |
| Report                          |     ✅     |
| Survey participant list         |     ✅     |
| Users                           |     ✅     |
| User-DAG assignment             |     ✅     |
| User Roles                      |     ✅     |
| User-Role assignment            |     ✅     |
| Version                         |     ✅     |

### Import

| Feature                         | Available |
|---------------------------------|:---------:|
| Arms                            |     ✅     |
| Data Access Groups              |     ✅     |
| Events                          |     ✅     |
| File                            |     ✅     |
| Instrument-event mapping        |     ✅     |
| Metadata                        |     ✅     |
| Records                         |     ✅     |
| Repeating instruments and events|     ✅     |
| Users                           |     ✅     |
| User-DAG assignment             |     ✅     |
| User Roles                      |     ✅     |
| User-Role assignment            |     ✅     |

### Delete

| Feature                         | Available |
|---------------------------------|:---------:|
| Arms                            |     ✅     |
| Data Access Groups              |     ✅     |
| Events                          |     ✅     |
| File                            |     ✅     |
| Records                         |     ✅     |
| Users                           |     ✅     |
| User Roles                      |     ✅     |

### Other

| Feature                        | Available |
|--------------------------------|:---------:|
| Generate next record name      |     ✅     |
| Switch data access group       |     ✅     |


## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/tkruer/go-redcap/tree/main/LICENSE) file for details.
