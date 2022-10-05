package main

import "github.com/go-digitalent-hacktiv8/session-6/2-gin/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
