package main

import (
	"log"
	server "pex/app"
)

func main() {
	err := server.RunServer()
	if err != nil {
		log.Fatal("Server error")
	}

}
