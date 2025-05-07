package problem

import (
	"container/list"
)

type TrieNode struct {
	children map[byte]*TrieNode
	failLink *TrieNode
	matches  []int
}

type PatternMatcher struct {
	root     *TrieNode
	patterns []string
}

func NewPatternMatcher(patterns []string) *PatternMatcher {
	matcher := &PatternMatcher{root: &TrieNode{children: make(map[byte]*TrieNode)}, patterns: patterns}
	for i, pattern := range patterns {
		matcher.insertPattern(pattern, i)
	}
	matcher.buildFailureLinks()
	return matcher
}

func (matcher *PatternMatcher) insertPattern(pattern string, index int) {
	current := matcher.root
	for i := 0; i < len(pattern); i++ {
		char := pattern[i]
		if current.children[char] == nil {
			current.children[char] = &TrieNode{children: make(map[byte]*TrieNode)}
		}
		current = current.children[char]
	}
	current.matches = append(current.matches, index)
}

func (matcher *PatternMatcher) buildFailureLinks() {
	queue := list.New()
	matcher.root.failLink = matcher.root

	for _, child := range matcher.root.children {
		child.failLink = matcher.root
		queue.PushBack(child)
	}

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		current := element.Value.(*TrieNode)

		for char, child := range current.children {
			failNode := current.failLink
			for failNode != matcher.root && failNode.children[char] == nil {
				failNode = failNode.failLink
			}
			if failNode.children[char] != nil {
				child.failLink = failNode.children[char]
			} else {
				child.failLink = matcher.root
			}
			child.matches = append(child.matches, child.failLink.matches...)
			queue.PushBack(child)
		}
	}
}

func (matcher *PatternMatcher) FindMatches(text string) []int {
	var result []int
	current := matcher.root

	for i := 0; i < len(text); i++ {
		char := text[i]
		for current != matcher.root && current.children[char] == nil {
			current = current.failLink
		}
		if current.children[char] != nil {
			current = current.children[char]
		}
		for _, patternIndex := range current.matches {
			startIndex := i - len(matcher.patterns[patternIndex]) + 1
			result = append(result, startIndex)
		}
	}
	return result
}
