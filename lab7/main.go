package main

import (
	"bufio"
	"lab7/problem"
	"os"
	"strconv"
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

	file, _ := os.Open("input2.dat")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var patterns []string
	for i := 0; i < n; i++ {
		scanner.Scan()
		patterns = append(patterns, scanner.Text())
	}

	matcher := problem.NewPatternMatcher(patterns)
	result := matcher.FindMatches(text)

	outputFile, _ := os.Create("output2.dat")
	defer outputFile.Close()
	output := make([]string, len(result))
	for i, val := range result {
		output[i] = strconv.Itoa(val)
	}
	outputFile.WriteString(strings.Join(output, " ") + "\n")
}
