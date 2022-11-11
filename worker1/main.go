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

	w1 := worker.New(c, "hello-world", wo)

	w1.RegisterWorkflow(e.MyWorkflow)
	w1.RegisterActivity(e.MyActivity)
	w1.RegisterWorkflow(e.MyWorkflow1)
	w1.RegisterActivity(e.MyActivity1)

	if err := w1.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
