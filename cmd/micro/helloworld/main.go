package main

import (
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

func main() {
	// create a new service
	service := micro.NewService(
		micro.Name("helloworld"),
		micro.Version("latest"),
	)

	// initialise flags
	service.Init()

	// start the service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
