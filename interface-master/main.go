package main

import (
	"DemoTools/interface-master/router"
	"log"
)

func main() {
	r := router.InitRouter()
	log.Fatal(r.Run(":8080"))
}
