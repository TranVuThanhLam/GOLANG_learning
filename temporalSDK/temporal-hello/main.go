package main

import (
	"context"
	"fmt"
	"log"
	"temporal-hello/app"

	"go.temporal.io/sdk/client"
)

func main() {
	// create a connect to to Temporal server
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Không thể tạo Temporal client:", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "say-hello-workflow",
		TaskQueue: "demo-task-queue",
	}

	param := app.GreetingParam{Name: "Lâm"}

	we, err := c.ExecuteWorkflow(context.Background(), options, app.SayHelloWorkflow, param)
	if err != nil {
		log.Fatalln("Không thể khởi chạy workflow:", err)
	}

	fmt.Println("🚀 Workflow đã chạy với ID:", we.GetID())

	var results []app.GreetingResult
	err = we.Get(context.Background(), &results)
	if err != nil {
		log.Fatalln("Lỗi khi lấy kết quả:", err)
	}

	for i, r := range results {
		fmt.Printf("📨 Kết quả %d: %s\n", i+1, r.Message)
	}
}
