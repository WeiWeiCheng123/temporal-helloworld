package main

import (
	"fmt"
	"log"

	e "eric/helloworld"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "eric_test",
	})
	if err != nil {
		fmt.Println("Unable to create client", err)
		panic(err)
	}
	defer c.Close()

	wo := worker.Options{
		MaxConcurrentActivityTaskPollers: 10,
		MaxConcurrentWorkflowTaskPollers: 10,
		OnFatalError: func(err error) {
			log.Fatalln("Fatal error", err)
		},
	}

	w2 := worker.New(c, "hello-world", wo)

	w2.RegisterWorkflow(e.MyWorkflow)
	w2.RegisterActivity(e.MyActivity)
	w2.RegisterWorkflow(e.MyWorkflow1)
	w2.RegisterActivity(e.MyActivity1)

	if err := w2.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Unable to start worker", err)
	}
	/*
		if err := w2.Start(); err != nil {
			log.Fatalln("Unable to start worker", err)
		}

		select {}
	*/
}
