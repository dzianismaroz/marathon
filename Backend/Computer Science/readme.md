Let’s explore common **CS data structures** and their examples in **Go**.

---

### **1. Stack**

A **stack** is a linear data structure that follows the **Last In, First Out (LIFO)** principle. This means that the last element added to the stack is the first one to be removed.

#### **Operations**:
- **Push**: Adds an element to the top of the stack.
- **Pop**: Removes the top element from the stack.
- **Peek/Top**: Returns the top element without removing it.
- **IsEmpty**: Checks whether the stack is empty.

#### **Go Example**:

```go
package main

import (
	"fmt"
)

// Stack struct
type Stack struct {
	items []int
}

// Push adds an item to the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes an item from the stack and returns it
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Peek returns the top item without removing it
func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := &Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// Pop an element
	item, _ := stack.Pop()
	fmt.Println("Popped item:", item)

	// Peek at the top element
	item, _ = stack.Peek()
	fmt.Println("Top item:", item)

	// Check if the stack is empty
	fmt.Println("Is stack empty?", stack.IsEmpty())
}
```

---

### **2. Heap**

A **heap** is a special tree-based data structure that satisfies the **heap property**:
- **Max-Heap**: The key at each node is greater than or equal to the keys of its children.
- **Min-Heap**: The key at each node is less than or equal to the keys of its children.

In Go, heaps can be implemented using the `container/heap` package.

#### **Go Example (Min-Heap)**:

```go
package main

import (
	"container/heap"
	"fmt"
)

// MinHeap implements heap.Interface and holds the data.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)     { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &MinHeap{7, 1, 3}
	heap.Init(h)

	heap.Push(h, 4)
	heap.Push(h, 2)

	fmt.Println("Min-Heap:", *h)

	// Pop elements
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
```

---

### **3. Queue**

A **queue** is a linear data structure that follows the **First In, First Out (FIFO)** principle. Elements are added to the rear and removed from the front.

#### **Operations**:
- **Enqueue**: Adds an element to the rear of the queue.
- **Dequeue**: Removes an element from the front of the queue.
- **Peek/Front**: Returns the front element without removing it.
- **IsEmpty**: Checks whether the queue is empty.

#### **Go Example**:

```go
package main

import "fmt"

// Queue struct
type Queue struct {
	items []int
}

// Enqueue adds an item to the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes an item from the front of the queue
func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Peek returns the front item without removing it
func (q *Queue) Peek() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	return q.items[0], true
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	queue := &Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	// Dequeue an element
	item, _ := queue.Dequeue()
	fmt.Println("Dequeued item:", item)

	// Peek at the front element
	item, _ = queue.Peek()
	fmt.Println("Front item:", item)

	// Check if the queue is empty
	fmt.Println("Is queue empty?", queue.IsEmpty())
}
```

---

### **4. Tree**

A **tree** is a hierarchical data structure with nodes connected by edges. A **binary tree** is a common type, where each node has at most two children.

#### **Operations**:
- **Insert**: Adds a node to the tree.
- **Search**: Finds an element in the tree.
- **Traversal**: Visit nodes in a specific order (Pre-order, In-order, Post-order).

#### **Go Example (Binary Tree)**:

```go
package main

import "fmt"

// TreeNode represents a node in the binary tree
type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// Insert inserts a new value into the binary search tree
func (node *TreeNode) Insert(value int) {
	if value < node.value {
		if node.left == nil {
			node.left = &TreeNode{value: value}
		} else {
			node.left.Insert(value)
		}
	} else {
		if node.right == nil {
			node.right = &TreeNode{value: value}
		} else {
			node.right.Insert(value)
		}
	}
}

// InOrderTraversal traverses the tree in-order (left, root, right)
func (node *TreeNode) InOrderTraversal() {
	if node != nil {
		node.left.InOrderTraversal()
		fmt.Println(node.value)
		node.right.InOrderTraversal()
	}
}

func main() {
	root := &TreeNode{value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)

	// In-order traversal
	fmt.Println("In-order traversal:")
	root.InOrderTraversal()
}
```

---

### **5. Map (HashMap)**

A **map** (also called a hash map) is a collection of key-value pairs. Go provides built-in support for maps, which allow fast retrieval and updates.

#### **Operations**:
- **Insert**: Adds a key-value pair.
- **Get**: Retrieves a value by key.
- **Delete**: Removes a key-value pair.
- **Exists**: Checks if a key exists.

#### **Go Example**:

```go
package main

import "fmt"

func main() {
	// Create a map
	m := make(map[string]int)

	// Insert key-value pairs
	m["apple"] = 5
	m["banana"] = 3

	// Get value by key
	appleCount := m["apple"]
	fmt.Println("Apple count:", appleCount)

	// Delete key-value pair
	delete(m, "banana")

	// Check if key exists
	_, exists := m["banana"]
	fmt.Println("Banana exists:", exists)

	// Iterate over map
	for key, value := range m {
		fmt.Println(key, ":", value)
	}
}
```

---

### **6. List**

A **list** (or **linked list**) is a linear data structure where each element (node) contains a value and a reference (link) to the next element in the sequence. There are **singly linked lists** and **doubly linked lists**.

#### **Operations**:
- **Insert**: Adds a new node.
- **Delete**: Removes a node.
- **Traverse**: Iterates through the list.

#### **Go Example (Singly Linked List)**:

```go
package main

import "fmt"

// ListNode represents a node in the linked list
type ListNode struct {
	value int
	next  *ListNode
}

// Append appends a new node to the list
func (node *ListNode) Append(value int) {
	newNode := &ListNode{value: value}
	for node.next != nil {
		node = node.next
	}
	node.next = newNode
}

// Traverse traverses the list
func (node *ListNode) Traverse() {
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

func main() {
	// Create a linked list with one node
	head := &ListNode{value: 10}

	// Append more nodes
	head.Append(20)
	head.Append(30)

	// Traverse the list
	head.Traverse()
}
```

---

### **Other Essential Data Structures**

1. **Graph**: A collection of nodes (vertices) connected by edges. Useful for representing networks (e.g., social networks, computer networks). Can be implemented using an adjacency matrix or adjacency list.

2. **Trie (Prefix Tree)**: A tree-like data structure that stores strings in a way that allows for fast search, insert, and prefix-based queries. Often used in autocomplete and dictionary applications.

---

### **Summary**

- **Stack**: LIFO structure, useful for undo operations, function calls (recursion).
- **Heap**: A tree-based structure used for efficient priority queue management.
- **Queue**: FIFO structure, used in breadth-first search and task scheduling.
- **Tree**: Hierarchical structure, useful in sorting and searching.
- **Map (HashMap)**: Key-value pairs, with fast lookups and updates.
- **List (Linked List)**: Dynamic size structure, with efficient insertions and deletions.

These data structures form the foundation of many algorithms, and understanding them is key to writing efficient and optimized code in Go.