package main

import (
	"log"

	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/Chapter08/backend/src/rest"
)

func main() {
	log.Println("Main log....")
	log.Fatal(rest.RunAPI("127.0.0.1:8000"))
}
