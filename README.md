# Go Hacks

A project to demonstrate the capabilities of the Go syntax, from basics to more complex things.

#### Sections
* [variables](./sections/variables/variables.go)
* [functions](./sections/functions/functions.go)
* [flow](./sections/flow/flow.go)
* [moretypes](./sections/moretypes/moretypes.go)

### Commands

The list of most commonly used commands are the following:

1. `go mod init`

    Initializes a new module in your current directory, creating a `go.mod` file.<br/>
    __Usage:__ `go mod init <module-name>`<br/>

2. `go mod tidy`

    Adds missing and removes unused modules.
    Ensures that the `go.mod` file matches the source code in the module.<br/>
    __Usage:__ `go mod tidy`

3. `go build`

    Compiles your Go program.
    By default, it does not produce an output file. Instead, it saves the compiled binary in a temporary location.
    To save the compiled binary, you can specify a name: `go build -o <binary-name>`.<br/>
    __Usage:__ `go build` or `go build -o myapp`


4. `go run`

    Compiles and runs your Go program.
    Useful for quick testing and development.<br/>
    Usage: `go run <file-name>`<br/>
    Example: `go run .` or `go run main.go`

5. `go test`

    Runs tests in the current package.
    Go identifies any file ending in `_test.go` as a test.<br/>
    __Usage:__ `go test`<br/>
    _To run tests in all subdirectories, use_ `go test ./...`.

6. `go get`

    Adds a dependency to your module and installs it.<br/>
    __Usage:__ `go get <package-name>`<br/>
    __Example:__ `go get github.com/gin-gonic/gin`

7. `go install`

    Compiles and installs a Go program (typically used for installing Go tools).
    The binary is saved to the bin directory in your `GOPATH`.<br/>
    __Usage:__ `go install <package-name>`

8. `go fmt`

    Formats your Go source code according to the Go standards.<br/>
    __Usage:__ `go fmt <file-name>`<br/>
    _To run format in the current directory and its subdirectories, use_ `go fmt ./..`.

9. `go vet`

    Analyzes your source code and reports potential errors like unreachable code, unnecessary assignments, etc.<br/>
    __Usage:__ `go vet <file-name>`<br/>
    _To run format for all files in the current directory and its subdirectories, use_ `go vet ./..`.

10. `go env`

    Prints Go environment information.
    Useful for debugging and understanding your Go setup.<br/>
    __Usage:__ `go env`
