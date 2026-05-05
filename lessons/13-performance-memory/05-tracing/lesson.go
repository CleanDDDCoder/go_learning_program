package tracing

import (
	"context"
	"os"
	"runtime/trace"
)

// StartTrace begins a trace and returns a function to stop it.
//
// TODO: Implement trace start using trace.Start.
// TODO: The returned function should call trace.Stop.
func StartTrace(filename string) (stop func() error) {
	f, err := os.Create(filename)
	if err != nil {
		return func() error { return err }
	}
	trace.Start(f)
	return func() error {
		trace.Stop()
		return f.Close()
	}
}

// TaskRegion marks a region in the trace for analysis.
//
// TODO: Implement task region using trace.WithRegion.
func TaskRegion(ctx context.Context, label string, fn func()) {
	trace.WithRegion(ctx, label, fn)
}

// Task creates a named task that can be viewed in the trace viewer.
//
// TODO: Implement task creation using trace.NewTask.
// Note: trace.NewTask returns (context.Context, *Task), not (*Task, error).
func Task(ctx context.Context, label string) (context.Context, *trace.Task) {
	return trace.NewTask(ctx, label)
}
