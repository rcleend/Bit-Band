package main

import (
	"fmt"
	"github.com/rcleend/bitband/sockhandler"
)

func main() {
	fmt.Println("Starting server...")
	sockhandler.StartConnection()
}
