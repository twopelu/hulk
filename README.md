# Hulk CLI

Hulk is a tool written in Go to make bulk changes in Git repos.

## Dependencies

Bitbucket API: https://pkg.go.dev/github.com/ktrysmt/go-bitbucket

## Installation

1. Move to lib folder and execute `go build`
2. Move to cli folder and execute `go build`
2. Move to gui folder and execute `go build`
3. In the root folder execute `go install`
4. Execute `hulk` to check it works

`go run main.py`

`go test` (TODO)

`go build`

## Usage

`hulk <user> <pass> <owner>`

ie. `hulk me@example.com 1234 me`