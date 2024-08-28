# Curseforge Fingerprint

A Go library to compute fingerprints for curseforge lookups

## Installation

To install the dependencies, run:

```sh
go mod tidy
```

## Usage

To use the library, import it and call the `GetFingerprintFor` function:

```go
package main

import (
    "fmt"
    "github.com/meza/curseforge-fingerprint-go"
)

func main() {
    filePath := "path/to/your/file"
    fingerprint := curseforgeFingerprint.GetFingerprintFor(filePath)
    fmt.Println("Fingerprint:", fingerprint)
}
```

## Testing

To run the tests, use the following command:

```sh
go test ./...
```

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

## License

This project is licensed under the MIT License.
