package main

import (
	"log"

	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/Chapter07/backend/src/rest"
)

func main() {
	log.Println("Main log....")
	log.Fatal(rest.RunAPI(":8000"))
}
