package main

import (
	"context"
	"fmt"
	"github.com/nanjingblue/go-easy-distribute/log"
	"github.com/nanjingblue/go-easy-distribute/registry"
	"github.com/nanjingblue/go-easy-distribute/service"
	stdlog "log"
)

func main() {
	log.Run("./destination.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName:      registry.LogService,
		ServiceURL:       serviceAddress,
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateURL: serviceAddress + "/services",
	}

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)
	if err != nil {
		stdlog.Fatalln(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service.")
}
