package main

import (
	"flag"
)

func main() {
	port := flag.Int("port", 1501, "port for the server to listen")
	flag.Parse()

	startServer(*port)
}
