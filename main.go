package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil{
		log.Fatal(err)
	}
	server := NewAPIServer(":3003")
	server.Run()
     fmt.Printf("%+v\n", store)
}
