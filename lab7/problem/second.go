package problem

type TrieNode struct {
	children    map[byte]*TrieNode
	isEndOfWord bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[byte]*TrieNode)}
}

func BuildPrefixTre(patterns []string) *TrieNode {
	root := NewTrieNode()
	for _, pattern := range patterns {
		current := root
		for i := 0; i < len(pattern); i++ {
			char := pattern[i]
			if current.children[char] == nil {
				current.children[char] = NewTrieNode()
			}
			current = current.children[char]
		}
		current.isEndOfWord = true
	}
	return root
}

func MatchPatternsInText(text string, patterns []string) []int {
	root := BuildPrefixTre(patterns)
	var matchIndices []int

	for i := 0; i < len(text); i++ {
		current := root
		position := i
		for position < len(text) && current.children[text[position]] != nil {
			current = current.children[text[position]]
			if current.isEndOfWord {
				matchIndices = append(matchIndices, i)
				break
			}
			position++
		}
	}
	return matchIndices
}

func ParseInt(s string) int {
	result := 0
	for _, ch := range s {
		result = result*10 + int(ch-'0')
	}
	return result
}
