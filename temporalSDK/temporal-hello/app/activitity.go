package app

import (
	"context"
	"fmt"
	"time"
)

type GreetingParam struct {
	Name string
}

type GreetingResult struct {
	Message string
}

func SayHelloActivity(ctx context.Context, param GreetingParam) (*GreetingResult, error) {
	time.Sleep(2 * time.Second) // giả lập hoạt động chậm
	msg := fmt.Sprintf("Xin chào %s từ Activity!", param.Name)
	return &GreetingResult{Message: msg}, nil
}
