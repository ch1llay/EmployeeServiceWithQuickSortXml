package main

import (
	"log"
)

func main() {

	if err := srv.Run(config, handlers); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
