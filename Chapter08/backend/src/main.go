package main

import (
	"log"

	"github.com/minaandrawos/Hands-On-Full-Stack-Development-with-Go/7-Testing-and-benchmarking/backend/src/rest"
)

func main() {
	log.Println("Main log....")
	log.Fatal(rest.RunAPI("127.0.0.1:8000"))
}
