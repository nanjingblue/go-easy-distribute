package main

import (
	"context"
	"fmt"
	"github.com/nanjingblue/go-easy-distribute/grades"
	"github.com/nanjingblue/go-easy-distribute/registry"
	"github.com/nanjingblue/go-easy-distribute/service"
	stdlog "log"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	r := registry.Registration{
		ServiceName: registry.GradingService,
		ServiceURL:  serviceAddress,
	}

	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers)
	if err != nil {
		stdlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
