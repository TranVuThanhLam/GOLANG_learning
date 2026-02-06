package main

import (
	"log"
	"temporal-hello/app"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Lỗi kết nối Temporal Client:", err)
	}
	defer c.Close()

	workerOptions := worker.Options{
		MaxConcurrentActivityExecutionSize: 3,
		WorkerActivitiesPerSecond:          10,
		EnableLoggingInReplay:              true,
		StickyScheduleToStartTimeout:       time.Second * 10,
		Identity:                           "demo-worker-lam",
	}

	w := worker.New(c, "demo-task-queue", workerOptions)

	w.RegisterWorkflow(app.SayHelloWorkflow)
	w.RegisterActivity(app.SayHelloActivity)

	log.Println("🛠️ Worker đang chạy...")
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Worker bị lỗi:", err)
	}
}
