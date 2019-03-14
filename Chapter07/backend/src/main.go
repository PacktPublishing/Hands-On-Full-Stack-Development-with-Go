package main

import (
	"log"

	"github.com/minaandrawos/Hands-On-Full-Stack-Development-with-Go/6-Advanced-Web/backend/src/rest"
)

func main() {
	log.Println("Main log....")
	log.Fatal(rest.RunAPI(":8000"))
}
