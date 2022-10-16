package main

import "log"

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
