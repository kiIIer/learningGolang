package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type question struct {
	q string
	a string
}

func main() {
	nameOfFile := flag.String("file", "questions.csv", "Name of questions file")
	// timeForAnswers := flag.Int("time", 30, "Time  given to answer")
	flag.Parse()

	file, err := os.Open(*nameOfFile)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	questions := parseLines(lines)

	score := 0
	for i, p := range questions {
		fmt.Printf("Solve #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			score++
		}
	}
}

func parseLines(lines [][]string) []question {
	ret := make([]question, len(lines))
	for i, line := range lines {
		ret[i] = question{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}
