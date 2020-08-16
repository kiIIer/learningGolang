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
	nameOfFile := flag.String("file", "questions.csv", "Name of questions file")
	timeForAnswers := flag.Int("time", 30, "Time  given to answer")
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
	chr := make(chan bool)
	cht := make(chan bool)
	go timeOuter(*timeForAnswers, cht)
	for i, p := range questions {
		go ranger(i, p, &score, chr)
		select {
		case <-chr:
			break
		case <-cht:
			exit(score)
			return
		}
	}
	exit(score)
	return
}

func timeOuter(timeOut int, cht chan<- bool) {
	time.Sleep(time.Duration(timeOut) * time.Second)
	cht <- true
}

func ranger(i int, p question, score *int, chr chan<- bool) {
	fmt.Printf("Solve #%d: %s = \n", i+1, p.q)
	var answer string
	fmt.Scanf("%s\n", &answer)
	if answer == p.a {
		*score++
	}
	chr <- true
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

func exit(score int) {
	fmt.Printf("You scored %v out of 12 \n \n", score)
	if score == 12 {
		fmt.Println("Well done!")
	}
}
