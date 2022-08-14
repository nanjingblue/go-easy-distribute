package main

import (
	"context"
	"fmt"
	"github.com/nanjingblue/go-easy-distribute/log"
	"github.com/nanjingblue/go-easy-distribute/service"
	stdlog "log"
)

func main() {
	log.Run("./destination.log")
	host, port := "localhost", "4000"
	ctx, err := service.Start(
		context.Background(),
		"Log service",
		host,
		port,
		log.RegisterHandlers,
	)
	if err != nil {
		stdlog.Fatalln(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service.")
}
