package stack_test

import (
	"reflect"
	"testing"

	stack "dzianismaroz.github.com/marathon/stack/pkg"
)

func TestStack(t *testing.T) {
	tests := []struct {
		name           string
		operations     func(*stack.Stack[int]) interface{}
		expectedResult interface{}
	}{
		{
			name: "Push one element and check length",
			operations: func(s *stack.Stack[int]) interface{} {
				s.Push(1)
				return s.Len()
			},
			expectedResult: 1,
		},
		{
			name: "Push and pop one element",
			operations: func(s *stack.Stack[int]) interface{} {
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
			operations: func(s *stack.Stack[int]) interface{} {
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
			operations: func(s *stack.Stack[int]) interface{} {
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
			operations: func(s *stack.Stack[int]) interface{} {
				return s.Len()
			},
			expectedResult: 0,
		},
		{
			name: "Push and pop multiple elements interleaved",
			operations: func(s *stack.Stack[int]) interface{} {
				s.Push(10)
				s.Push(20)
				_, _ = s.Pop()
				s.Push(30)
				elem, _ := s.Pop()
				return elem
			},
			expectedResult: 30, // The last pushed element
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := stack.NewStack[int]()
			result := tt.operations(stack)

			// Compare results properly for slices or other types
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}
