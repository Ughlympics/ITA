package problem

import (
	"fmt"
)

type Node struct {
	ID       int
	Children map[byte]*Node
}

var nextID int
var edges []string

func insert(root *Node, pattern string) {
	current := root
	for i := 0; i < len(pattern); i++ {
		c := pattern[i]
		if _, exists := current.Children[c]; !exists {
			newNode := &Node{ID: nextID, Children: make(map[byte]*Node)}
			nextID++
			current.Children[c] = newNode
			edges = append(edges, fmt.Sprintf("%d->%d:%c", current.ID, newNode.ID, c))
		}
		current = current.Children[c]
	}
}

func BuildPrefixTree(patterns []string) []string {
	root := &Node{ID: 0, Children: make(map[byte]*Node)}
	nextID = 1
	edges = []string{}

	for _, pattern := range patterns {
		insert(root, pattern)
	}

	return edges
}
