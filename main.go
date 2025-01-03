package main

import (
	"github.com/go-to/egp-backend/router"
)

func main() {
	port := 8080
	router.New(port)
}
