package main

import (
	"fmt"
	"lab7/problem"
	"os"
	"strings"
)

func main() {
	// file, err := os.Open("input.dat")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// scanner.Scan()
	// var n int
	// fmt.Sscanf(scanner.Text(), "%d", &n)

	// var patterns []string
	// for scanner.Scan() {
	// 	line := strings.TrimSpace(scanner.Text())
	// 	if len(line) > 0 {
	// 		patterns = append(patterns, line)
	// 	}
	// }

	// output := problem.BuildPrefixTree(patterns)

	// outFile, err := os.Create("output.dat")
	// if err != nil {
	// 	panic(err)
	// }
	// defer outFile.Close()

	// for _, line := range output {
	// 	outFile.WriteString(line + "\n")
	// }

	data, err := os.ReadFile("input2.dat")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	text := lines[0]
	patternCount := problem.ParseInt(lines[1])
	patterns := lines[2 : 2+patternCount]

	indices := problem.MatchPatternsInText(text, patterns)

	var output []string
	for _, idx := range indices {
		output = append(output, fmt.Sprintf("%d", idx))
	}

	err = os.WriteFile("output2.dat", []byte(strings.Join(output, " ")), 0644)
	if err != nil {
		panic(err)
	}
}
