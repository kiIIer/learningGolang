package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type question struct {
	q string
	a string
}

func main() {
	fileName := flag.String("file", "questions.csv", "Name of questions file")
	timeoutSeconds := flag.Int("timeout", 30, "Timeout (seconds) for all answers")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	questions := parseLines(lines)

	score := 0
	timeout := newTimeout(*timeoutSeconds)

quastionsLoop:
	for i, q := range questions {
		select {
		case answer := <-askUser(i+1, q.q):
			if answer == q.a {
				score++
			}
			break
		case <-timeout:
			fmt.Printf("\ntimeout of %v seconds expired.\n", *timeoutSeconds)
			break quastionsLoop
		}
	}

	printScore(score, len(lines))
}

func newTimeout(timeoutSeconds int) <-chan time.Time {
	return time.After(time.Duration(timeoutSeconds) * time.Second)
}

func askUser(questionNumber int, question string) <-chan string {
	fmt.Printf("Solve #%d: %s = ", questionNumber, question)
	answer := make(chan string)
	go func() {
		var text string
		fmt.Scanf("%s\n", &text)
		answer <- text
		close(answer)
	}()
	return answer
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

func printScore(score, total int) {
	fmt.Printf("You scored %v out of %v \n \n", score, total)
	if score == total {
		fmt.Println("Well done!")
	}
}
