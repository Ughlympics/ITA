package main

import (
	"bufio"
	"fmt"
	"lab7/problem"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.dat")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var n int
	fmt.Sscanf(scanner.Text(), "%d", &n)

	var patterns []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			patterns = append(patterns, line)
		}
	}

	output := problem.BuildPrefixTree(patterns)

	outFile, err := os.Create("output.dat")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	for _, line := range output {
		outFile.WriteString(line + "\n")
	}
}
