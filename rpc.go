package main

import (
	"net/rpc"
	"os"

	"github.com/aws/aws-lambda-go/lambda/messages"
)

// Invoke makes the request to the local lambda function running
// on the address specified.
func Invoke(addr, eventFile string, data []byte) {
	request := messages.InvokeRequest{Payload: data}
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(-2)
	}

	// Read event file
	if eventFile != "" {
		f, err := os.Open(eventFile)
		if err != nil {
			os.Stderr.WriteString("error opening file: " + err.Error() + "\n")
			os.Exit(-3)
		}

		fileInfo, _ := f.Stat()
		payload := make([]byte, fileInfo.Size())
		n, err := f.Read(payload)
		if int64(n) != fileInfo.Size() {
			os.Stderr.WriteString("error: could not read whole file" + "\n")
			os.Exit(-4)
		}
		if err != nil {
			os.Stderr.WriteString("error reading file: " + err.Error() + "\n")
			os.Exit(-5)
		}
		request.Payload = payload
	}

	var reply messages.InvokeResponse
	err = client.Call("Function.Invoke", request, &reply)
	if err != nil {
		os.Stderr.WriteString("call error: " + err.Error() + "\n")
		os.Exit(-6)
	}

	if reply.Error != nil {
		os.Stderr.WriteString("lambda returned error:\n")
		os.Stdout.WriteString(reply.Error.Message + "\n")
	} else {
		os.Stdout.WriteString(string(reply.Payload) + "\n")
	}
	os.Exit(0)
}
