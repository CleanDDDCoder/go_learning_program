package context_cancellation

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestCollectUntilDoneReadsClosedChannel(t *testing.T) {
	input := make(chan int, 3)
	input <- 1
	input <- 2
	input <- 3
	close(input)

	got := CollectUntilDone(context.Background(), input)
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("CollectUntilDone() = %v, want %v", got, want)
	}
}

func TestCollectUntilDoneStopsOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	input := make(chan int)
	done := make(chan []int, 1)

	go func() {
		done <- CollectUntilDone(ctx, input)
	}()

	input <- 1
	cancel()

	select {
	case got := <-done:
		if len(got) != 1 || got[0] != 1 {
			t.Fatalf("CollectUntilDone after cancel = %v, want [1]", got)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("CollectUntilDone did not return after cancellation")
	}
}
