package queue_test

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/dzianismaroz/marathon/queue/queue"
	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		name     string
		scenario func(*testing.T)
	}{
		{
			name: "should create empty queue",
			scenario: func(t *testing.T) {
				q := queue.New[int]()
				if q == nil {
					t.Error("queue should not be nil")
				}
			},
		},
		{
			name: "should push and pop elements",
			scenario: func(t *testing.T) {
				q := queue.New[int]()
				q.Push(1)
				q.Push(2)
				q.Push(3)

				val, ok := q.Pop()
				if !ok || val != 1 {
					t.Errorf("expected 3, got %v", val)
				}

				val, ok = q.Pop()
				if !ok || val != 2 {
					t.Errorf("expected 2, got %v", val)
				}

				val, ok = q.Pop()
				if !ok || val != 3 {
					t.Errorf("expected 1, got %v", val)
				}
			},
		},
		{
			name: "should handle empty queue operations",
			scenario: func(t *testing.T) {
				q := queue.New[string]()

				val, ok := q.Pop()
				if ok || val != "" {
					t.Error("pop from empty queue should return zero value and false")
				}

				val, ok = q.Peek()
				if ok || val != "" {
					t.Error("peek from empty queue should return zero value and false")
				}

				if !q.IsEmpty() {
					t.Error("empty queue should have length 0")
				}
			},
		},
		{
			name: "should peek without removing elements",
			scenario: func(t *testing.T) {
				q := queue.New[float64]()
				q.Push(1.1)
				q.Push(2.2)

				val, ok := q.Peek()
				if !ok || val != 1.1 {
					t.Errorf("expected 2.2, got %v", val)
				}

				if q.Size() != 2 {
					t.Errorf("expected length 2, got %d", q.Size())
				}
			},
		},
		{
			name: "should maintain LIFO order",
			scenario: func(t *testing.T) {
				q := queue.New[int]()
				values := []int{1, 2, 3, 4, 5}

				for _, v := range values {
					q.Push(v)
				}

				for i := 0; i <= len(values)-1; i++ {
					val, ok := q.Pop()
					if !ok || val != values[i] {
						t.Errorf("expected %d, got %d", values[i], val)
					}
				}
			},
		},
		{
			name: "Push and popAll elements",
			scenario: func(t *testing.T) {
				q := queue.New[int]()
				q.Push(10)
				q.Push(20)
				q.Push(30)
				q.Push(40)
				got := q.PopAll()
				expected := []int{10, 20, 30, 40}
				require.Empty(t, q.Size())
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("expected %d, got %d", expected, got)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.scenario)
	}
}

func BenchmarkQueue_Push(b *testing.B) {
	q := queue.New[int]()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}

func BenchmarkQueue_Pop(b *testing.B) {
	q := queue.New[int]()
	// Pre-fill the stack with N elements
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkQueue_Peek(b *testing.B) {
	q := queue.New[int]()

	q.Push(rand.Int())
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = q.Peek()
	}
}

func BenchmarkQueue_Size(b *testing.B) {
	q := queue.New[int]()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = q.Size()
	}
}

func BenchmarkQueue_PopAll(b *testing.B) {
	q := queue.New[int]()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = q.PopAll()
	}
}

func BenchmarkQueue_IsEmpty(b *testing.B) {
	q := queue.New[int]()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = q.IsEmpty()
	}
}

func BenchmarkQueue_Clear(b *testing.B) {
	q := queue.New[int]()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Clear()
	}
}
