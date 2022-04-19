package main

import (
	"github.com/tellmesomuch/go-api/pkg/server"
)

func main() {
	server.Serve("0.0.0.0:80")
}
