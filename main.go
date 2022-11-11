package helloworld

import (
	"bufio"
	"log"
	"os"
	"time"

	"go.temporal.io/sdk/workflow"
)

func MyWorkflow(ctx workflow.Context) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	lines, err := readLines("test.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for i := 0; i < 1000; i++ {
		err := workflow.ExecuteActivity(ctx, MyActivity, lines[i]).Get(ctx, nil)
		if err != nil {
			return "", err
		}
	}

	return "Workflow completed!", nil
}

func MyWorkflow1(ctx workflow.Context) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	lines, err := readLines("test.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for i := 0; i < 100; i++ {
		err := workflow.ExecuteActivity(ctx, MyActivity1, lines[i]).Get(ctx, nil)
		if err != nil {
			return "", err
		}
	}

	return "Workflow1 completed!", nil
}
func MyActivity(input string) (string, error) {
	//time.Sleep(time.Second * 1)
	return "Hello " + input + "!", nil
}

func MyActivity1(input string) (string, error) {
	//time.Sleep(time.Second * 1)
	return "Hi " + input + "!", nil
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
