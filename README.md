
# Sectoid Go Kit

The Sectoid Go Kit is a collection of micro libraries and useful modules for Go development.

## Package Summaries

- **[ctxutils](./ctxutils/README.md)**: Utilities for managing context and closing resources in a safe and error-handling manner.
- **[envutils](./envutils/README.md)**: Utility functions for retrieving environment variables with default values and type conversions.
- **[grpcutils](./grpcutils/README.md)**: Utility functions for converting gRPC errors into corresponding HTTP status codes.
- **[iterables](./iterables/README.md)**: Utility functions for working with lists and maps.
- **[logmesh](./logmesh/README.md)**: Interfaces and implementations for logging with various log levels and methods.
- **[misc](./misc/README.md)**: Utility functions for various common tasks such as checking for nil pointers and retrying operations with timeouts.
- **[strutils](./strutils/README.md)**: Utility functions for string conversions and manipulations.
- **[supermath](./supermath/README.md)**: Utility functions for mathematical operations, including truncating floating-point numbers to a specific number of decimal places.
- **[timespace](./timespace/README.md)**: Utility functions for converting between various timestamp formats and time representations.

## Adding to Your Project

To use the Sectoid Go Kit in your Go project, follow these steps:

1. Initialize a new Go module (if you haven't already):
    ```sh
    go mod init your_project_name
    ```
2. Add the Sectoid Go Kit as a dependency:
    ```sh
    go get github.com/Sectoid-Systems/sectoid-go-kit
    ```

3. Import the desired packages in your Go files:
    ```go
    import (
        "github.com/Sectoid-Systems/sectoid-go-kit/ctxutils"
        "github.com/Sectoid-Systems/sectoid-go-kit/envutils"
        // import other packages as needed
    )
    ```

## Development Instructions

To contribute to the Sectoid Go Kit, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/Sectoid-Systems/sectoid-go-kit.git
    ```
2. Navigate to the project directory:
    ```sh
    cd sectoid-go-kit
    ```
3. Install the required dependencies:
    ```sh
    go mod tidy
    ```
4. Make your changes and run the tests:
    ```sh
    go test ./...
    ```

## Makefile Commands

The project uses a Makefile to automate common tasks. Below is a summary of the available commands:

- `make build`: Compiles the project.
- `make test`: Runs the tests.
- `make clean`: Cleans the build artifacts.
- `make lint`: Runs the linter to check for code style issues.
- `make fmt`: Formats the code.

To use these commands, simply run `make <command>` in your terminal.

## License

The Sectoid Go Kit is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more information.
