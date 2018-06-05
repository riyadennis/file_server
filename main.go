package main

import (
	"github.com/file_server/handlers"
	"flag"
)

func main() {
	port := flag.String("port", "8080", "Port for web server")
	flag.Parse()
	handlers.Run(*port)
}
