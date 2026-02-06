package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func SayHelloWorkflow(ctx workflow.Context, param GreetingParam) ([]GreetingResult, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Gọi nhiều activity song song
	names := []string{"Lâm", "Trang", "Minh"}
	futures := make([]workflow.Future, len(names))
	results := make([]GreetingResult, len(names))

	for i, name := range names {
		p := GreetingParam{Name: name}
		futures[i] = workflow.ExecuteActivity(ctx, SayHelloActivity, p)
	}

	// Lấy kết quả khi sẵn sàng
	for i, future := range futures {
		if future.IsReady() {
			err := future.Get(ctx, &results[i])
			if err != nil {
				return nil, err
			}
		} else {
			// Nếu chưa sẵn sàng, vẫn Get (nó sẽ block cho đến khi xong)
			err := future.Get(ctx, &results[i])
			if err != nil {
				return nil, err
			}
		}
	}

	return results, nil
}
