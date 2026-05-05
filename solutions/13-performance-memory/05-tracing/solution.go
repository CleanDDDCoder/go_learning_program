package tracing

import (
	"context"
	"os"
	"runtime/trace"
)

// StartTraceFixed begins a trace and returns a function to stop it.
func StartTraceFixed(filename string) (stop func() error) {
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

// TaskRegionFixed marks a region in the trace for analysis.
func TaskRegionFixed(ctx context.Context, label string, fn func()) {
	trace.WithRegion(ctx, label, fn)
}

// TaskFixed creates a named task that can be viewed in the trace viewer.
func TaskFixed(ctx context.Context, label string) (context.Context, *trace.Task) {
	return trace.NewTask(ctx, label)
}