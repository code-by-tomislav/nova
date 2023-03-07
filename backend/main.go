package main

import (
	"log"
	"nova/utils"
	"nova/web"
)

func main() {
	log.Println("Starting server...")

	c := utils.Configure()
	web.Server(c)
}
