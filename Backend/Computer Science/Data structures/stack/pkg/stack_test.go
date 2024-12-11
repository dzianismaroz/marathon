package stack_test

import (
	"reflect"
	"testing"

	stack "dzianismaroz.github.com/marathon/stack/pkg"
)

func TestStack(t *testing.T) {
	tests := []struct {
		name           string
		operations     func(*stack.Stack[int]) any
		expectedResult any
	}{
		{
			name: "Push one element and check length",
			operations: func(s *stack.Stack[int]) any {
				s.Push(1)
				return s.Size()
			},
			expectedResult: uint(1),
		},
		{
			name: "Push and pop one element",
			operations: func(s *stack.Stack[int]) any {
				s.Push(42)
				elem, ok := s.Pop()
				return struct {
					Element int
					Ok      bool
				}{elem, ok}
			},
			expectedResult: struct {
				Element int
				Ok      bool
			}{42, true},
		},
		{
			name: "Pop from empty stack",
			operations: func(s *stack.Stack[int]) any {
				elem, ok := s.Pop()
				return struct {
					Element int
					Ok      bool
				}{elem, ok}
			},
			expectedResult: struct {
				Element int
				Ok      bool
			}{0, false}, // Default int value is 0
		},
		{
			name: "Push multiple elements and check order",
			operations: func(s *stack.Stack[int]) any {
				s.Push(1)
				s.Push(2)
				s.Push(3)
				results := []int{}
				for {
					elem, ok := s.Pop()
					if !ok {
						break
					}
					results = append(results, elem)
				}
				return results
			},
			expectedResult: []int{3, 2, 1}, // Stack follows LIFO order
		},
		{
			name: "Check length on empty stack",
			operations: func(s *stack.Stack[int]) any {
				return s.Size()
			},
			expectedResult: uint(0),
		},
		{
			name: "Push and pop multiple elements interleaved",
			operations: func(s *stack.Stack[int]) any {
				s.Push(10)
				s.Push(20)
				_, _ = s.Pop()
				s.Push(30)
				elem, _ := s.Pop()
				return elem
			},
			expectedResult: 30, // The last pushed element
		},

		{
			name: "Push and popAll elements",
			operations: func(s *stack.Stack[int]) any {
				s.Push(10)
				s.Push(20)
				s.Push(30)
				s.Push(40)
				res := s.PopAll()
				return res
			},
			expectedResult: []int{40, 30, 20, 10},
		},

		{
			name: "Push and Clear Stack",
			operations: func(s *stack.Stack[int]) any {
				s.Push(10)
				s.Push(20)
				s.Push(30)
				s.Push(40)
				s.Clear()
				return s.IsEmpty()
			},
			expectedResult: true, // stack must be empty after cleanup.
		},

		{
			name: "Push and Peek",
			operations: func(s *stack.Stack[int]) any {
				s.Push(1999)
				s.Push(1843)
				s.Push(2024)
				res, _ := s.Peek()
				return struct {
					Element int
					Size    uint
				}{res, s.Size()}
			},
			expectedResult: struct {
				Element int
				Size    uint
			}{2024, 3}, // stack must have full size as after Push.
		},

		{
			name: "Push on large ammount of data",
			operations: func(s *stack.Stack[int]) any {
				for i := range 100_000 {
					s.Push(i)
				}
				res, _ := s.Peek()
				return struct {
					Element int
					Size    uint
				}{res, s.Size()}
			},
			expectedResult: struct {
				Element int
				Size    uint
			}{99_999, 100_000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := stack.New[int]()
			result := tt.operations(stack)

			// Compare results properly for slices or other types
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}
