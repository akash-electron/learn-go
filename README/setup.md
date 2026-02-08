# Go Setup Guide

Follow these steps to set up your environment for Go development.

## 1. Installation

Download and install Go from the official portal:

- **Website**: [golang.org/dl](https://golang.org/dl/)
- **Documentation**: [go.dev/doc](https://go.dev/doc/install)

Check your installation:

```bash
go version
```

## 2. Environment Variables

Ensure your `GOPATH` and `GOBIN` are correctly set up in your shell profile (`.zshrc` or `.bash_profile`):

```bash
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

## 3. Creating a Module

Go uses modules to manage dependencies. To start a new project:

```bash
mkdir my-project
cd my-project
go mod init github.com/yourusername/my-project
```

## 4. Useful Commands

| Command            | Description                             |
| ------------------ | --------------------------------------- |
| `go run main.go`   | Compiles and runs the code              |
| `go build`         | Compiles the package but doesn't run it |
| `go test ./...`    | Runs all tests in the current directory |
| `go fmt ./...`     | Formats all Go files in the project     |
| `go get <package>` | Downloads a new package                 |
| `go mod tidy`      | Cleans up unused dependencies           |
