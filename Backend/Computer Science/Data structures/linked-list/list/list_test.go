package list_test

import (
	"reflect"
	"testing"

	"github.com/dzianismaroz/marathon/linked-list/list"
)

func TestRemoveAt(t *testing.T) {
	tables := []struct {
		name     string
		list     []int
		index    uint
		expected []int
		wantErr  bool
	}{
		{
			name:     "remove from the beginning",
			list:     []int{1, 2, 3},
			index:    0,
			expected: []int{2, 3},
		},
		{
			name:     "remove from the middle",
			list:     []int{1, 2, 3},
			index:    1,
			expected: []int{1, 3},
		},
		{
			name:     "remove from the end",
			list:     []int{1, 2, 3},
			index:    2,
			expected: []int{1, 2},
		},
		{
			name:     "remove from empty list",
			list:     []int{},
			index:    0,
			expected: []int{},
			wantErr:  true,
		},
	}

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			t.Parallel()

			sut := list.BuildFrom(table.list)

			_, err := sut.RemoveAt(table.index)
			if err != nil && !table.wantErr {
				t.Errorf("expected to remove %d at index %d without error but got %v", table.list[table.index], table.index, err)
			}

			if err == nil && table.wantErr {
				t.Errorf("expected to remove %d at index %d with error but got nil", table.list[table.index], table.index)
			}

			if !reflect.DeepEqual(table.expected, sut.Items()) {
				t.Errorf("expected =%v but got = %v", table.expected, sut.Items())
			}
		})
	}
}

func TestInsertAt(t *testing.T) {
	tables := []struct {
		name     string
		input    []int
		idx      uint
		item     int
		expected []int
	}{
		{
			name:     "insert at the beginning",
			input:    []int{1, 2, 3},
			idx:      0,
			item:     4,
			expected: []int{4, 1, 2, 3},
		},
		{
			name:     "insert at the middle",
			input:    []int{1, 2, 3},
			idx:      1,
			item:     4,
			expected: []int{1, 4, 2, 3},
		},
		{
			name:     "insert at the end",
			input:    []int{1, 2, 3},
			idx:      3,
			item:     4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "insert into empty list",
			input:    []int{},
			idx:      0,
			item:     4,
			expected: []int{4},
		},
	}

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			t.Parallel()

			sut := list.BuildFrom(table.input)

			err := sut.InsertAt(table.idx, table.item)
			if err != nil {
				t.Fatalf("expected to insert %d at index %d without error but got %v", table.item, table.idx, err)
			}

			if !reflect.DeepEqual(table.expected, sut.Items()) {
				t.Errorf("expected =%v but got = %v", table.expected, sut.Items())
			}
		})
	}
}

func TestIndexOf(t *testing.T) {
	tables := []struct {
		name     string
		list     []int
		item     int
		expected uint
		found    bool
	}{
		{
			name:     "item present at the beginning",
			list:     []int{1, 2, 3},
			item:     1,
			expected: 0,
			found:    true,
		},
		{
			name:     "item present in the middle",
			list:     []int{1, 2, 3},
			item:     2,
			expected: 1,
			found:    true,
		},
		{
			name:     "item present at the end",
			list:     []int{1, 2, 3},
			item:     3,
			expected: 2,
			found:    true,
		},
		{
			name:     "item not present",
			list:     []int{1, 2, 3},
			item:     4,
			expected: 0,
			found:    false,
		},
		{
			name:     "empty list",
			list:     []int{},
			item:     1,
			expected: 0,
			found:    false,
		},
	}

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			sut := list.BuildFrom(table.list)

			index, found := sut.IndexOf(table.item)
			if found != table.found || (found && index != table.expected) {
				t.Errorf("expected index %d and found %v, but got index %d and found %v", table.expected, table.found, index, found)
			}
		})
	}
}

func TestItems(t *testing.T) {
	tables := []struct {
		name     string
		list     []int
		expected []int
	}{
		{
			name:     "empty list",
			list:     []int{},
			expected: []int{},
		},
		{
			name:     "single element list",
			list:     []int{1},
			expected: []int{1},
		},
		{
			name:     "multi element list",
			list:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			t.Parallel()

			list := list.BuildFrom(table.list)

			result := list.Items()

			if !reflect.DeepEqual(result, table.expected) {
				t.Errorf("expected to get %v but got %v", table.expected, result)
			}
		})
	}
}

func TestSize(t *testing.T) {
	tables := []struct {
		name     string
		list     []int
		expected uint
	}{
		{
			name:     "empty list",
			list:     []int{},
			expected: 0,
		},
		{
			name:     "single element list",
			list:     []int{1},
			expected: 1,
		},
		{
			name:     "multiple elements list",
			list:     []int{1, 2, 3},
			expected: 3,
		},
	}

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			sut := list.BuildFrom(table.list)
			size := sut.Size()
			if size != table.expected {
				t.Errorf("expected size %d, but got %d", table.expected, size)
			}
		})
	}
}

func TestFirst(t *testing.T) {
	tables := []struct {
		name     string
		input    []int
		expected int
		ok       bool
	}{
		{
			name:     "first item from non-empty list",
			input:    []int{1, 2, 3},
			expected: 1,
			ok:       true,
		},
		{
			name:     "first item from empty list",
			input:    []int{},
			expected: 0,
		},
	}

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			sut := list.BuildFrom(table.input)

			actual, ok := sut.First()
			if !ok && table.ok {
				t.Errorf("expected to get first item from list %v but got false", table.input)
			}

			if ok && !table.ok {
				t.Errorf("expected to not get first item from list %v but got true", table.input)
			}

			if actual != table.expected {
				t.Errorf("expected first item from list %v to be %d but got %d", table.input, table.expected, actual)
			}
		})
	}
}

// ======================== BENCHMARKING ========================
func BenchmarkAppend(b *testing.B) {
	list := list.New[int]() // Create a new linked list
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Append(i) // Benchmark appending an element to the list
	}
}

func BenchmarkRemoveAt(b *testing.B) {
	list := list.BuildFrom([]int{1, 2, 3, 4, 5}) // Create a list with initial elements
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.RemoveAt(0) // Benchmark removing an element from the list
	}
}

func BenchmarkIndexOf(b *testing.B) {
	list := list.BuildFrom([]int{1, 2, 3, 4, 5}) // Create a list with initial elements
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, ok := list.IndexOf(3) // Benchmark finding the index of an element
		if !ok {
			b.Errorf("element not found")
		}
	}
}

func BenchmarkInsertAt(b *testing.B) {
	list := list.New[int]() // Create a new linked list
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.InsertAt(1, 4) // Benchmark inserting an element at a specific index
	}
}
