# awslambdarpc

[![Go Report Card](https://goreportcard.com/badge/github.com/blmayer/awslambdarpc)](https://goreportcard.com/report/github.com/blmayer/awslambdarpc)
[![Go Reference](https://pkg.go.dev/badge/github.com/blmayer/awslambdarpc.svg)](https://pkg.go.dev/github.com/blmayer/awslambdarpc)

> Small utility to make RPC requests to a locally running AWS Lambda, for development purposes.

## Installing

Run `go get github.com/blmayer/awslambdarpc`.

## Using

You need a running lambda, let's say in your computer port 3000, then to make a request to it,
run:

```awslambdarpc -a localhost:3000 -d '{"body": "Hello World!"}'```.

### Options

You can specify an input event to your lambda in 2 ways:

- Pointing to a JSON file
- Writing the JSON as an argument

For pointing to a file use the `-e` or `--event` option, e.g.:

```awslambdarpc -a localhost:3000 -e example.json```,

and passing the input directly is done with the `-d` or `--data` option, such as:

```awslambdarpc -a localhost:3000 -d '{"body": "Hello World!"}'```.

There is also the `-h` or `--help` flags that will give you further explanation.

## Why?

I couldn't setup a debugger using go and aws-sam-cli, so this way I could just compile the binary
for my function, attach the debugger on it and run this utility.
