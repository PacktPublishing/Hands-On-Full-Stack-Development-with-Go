package main

import (
	"log"

	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/final-application/src/rest"
)

func main() {
	log.Println("Main log....")
	log.Fatal(rest.RunAPI("127.0.0.1:8080"))
}
