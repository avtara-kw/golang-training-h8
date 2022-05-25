package main

import (
	"log"
	"sesi6-gin/routers"
)

func main() {
	port := ":4444"

	log.Default().Println("Server running at port", port)
	routers.StartServer().Run(port)
}
