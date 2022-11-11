package main

import (
	"context"
	e "eric/helloworld"
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
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

	workflowOptions := client.StartWorkflowOptions{
		ID:        "hello_world" + uuid.NewString(),
		TaskQueue: "hello-world",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, e.MyWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
		panic(err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result string
	if err := we.Get(context.Background(), &result); err != nil {
		log.Fatalln("Unable get workflow result", err)
	}

	log.Println("Workflow result:", result)
}
