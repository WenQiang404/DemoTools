package main

import (
	"interface/router"
	"log"
)

func main() {
	r := router.InitRouter()
	log.Fatal(r.Run(":8080"))
}
