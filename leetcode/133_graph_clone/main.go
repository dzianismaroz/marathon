package main

func main() {}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[int]*Node, len(node.Neighbors))

	return cloneNodes(visited, node)
}

func cloneNodes(visited map[int]*Node, node *Node) *Node {
	if node == nil {
		return nil
	}

	if n, ok := visited[node.Val]; ok {
		return n
	}

	newNode := &Node{Val: node.Val, Neighbors: make([]*Node, 0, len(node.Neighbors))}

	visited[node.Val] = newNode

	for _, neighbor := range node.Neighbors {
		others := cloneNodes(visited, neighbor)
		newNode.Neighbors = append(newNode.Neighbors, others)
	}

	return newNode
}
