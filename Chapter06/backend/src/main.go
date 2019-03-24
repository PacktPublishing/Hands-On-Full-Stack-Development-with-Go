package main

import (
	"log"

	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/Chapter06/backend/src/rest"
)

func main() {
	log.Println("Main log....")
	rest.RunAPI(":9090")
}
