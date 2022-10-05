package main

import "github.com/masred/scalable-web-service-with-golang/session-06/2-gin/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
