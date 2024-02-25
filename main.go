package main

import (
	"fmt"
)

func main() {
	server := NewAPIServer(":3003")
	server.Run()
	fmt.Println("Hi there!")
}
