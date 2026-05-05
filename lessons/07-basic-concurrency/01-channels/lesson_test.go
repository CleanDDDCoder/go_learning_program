package channels

import (
	"testing"
)

func TestSumConcurrent(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	result := SumConcurrent(ch)
	want := 55 // 1+2+3+...+10
	if result != want {
		t.Errorf("SumConcurrent() = %d, want %d", result, want)
	}
}

func TestSendNumbers(t *testing.T) {
	ch := make(chan int)
	go SendNumbers(ch, 5)

	expected := []int{1, 2, 3, 4, 5}
	for i, want := range expected {
		got, ok := <-ch
		if !ok {
			t.Errorf("channel closed early at index %d", i)
			break
		}
		if got != want {
			t.Errorf("received %d, want %d", got, want)
		}
	}
}

func TestConcurrentSumWithSendNumbers(t *testing.T) {
	ch := make(chan int)
	go SendNumbers(ch, 100)

	result := SumConcurrent(ch)
	want := 5050 // sum of 1..100
	if result != want {
		t.Errorf("SumConcurrent() = %d, want %d", result, want)
	}
}

func TestChannelClosedBehavior(t *testing.T) {
	ch := make(chan int)
	close(ch)

	result := SumConcurrent(ch)
	if result != 0 {
		t.Errorf("SumConcurrent() on closed empty channel = %d, want 0", result)
	}
}