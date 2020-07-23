package awslambdarpc

import (
	"os"
)

const help = `awslambdarpc is an utility to make requests to your local AWS Lambda
Usage:
  awslambdarpc [options]
Available options:
  -a
  --address	the address of your local running function, defaults to localhost:8080
  -e
  --event	path to the event JSON to be used as input
  -d
  --data	data passed to the function as input, in JSON format, defaults to "{}"
  -h
  --help	show this help
Examples:
  awslambdarpc -h localhost:3000 -e events/input.json
  awslambdarpc -h localhost:3000 -d '{"body": "Hello World!"}'`

func main() {
	addr := "localhost:8080"
	payload := []byte("{}")
	var eventFile string

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-a":
			fallthrough
		case "--address":
			i++
			addr = os.Args[i]
		case "-e":
			fallthrough
		case "--event":
			i++
			eventFile = os.Args[i]
		case "-d":
			fallthrough
		case "--data":
			i++
			payload = []byte(os.Args[i])
		case "-h":
			fallthrough
		case "--help":
			println(help)
			os.Exit(0)
		default:
			os.Stderr.WriteString("error: wrong argument\n")
			println(help)
			os.Exit(-1)
		}
	}

	Invoke(addr, eventFile, payload)
}
