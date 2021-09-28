package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	lines, err := readFile(csvFilename)
	if err != nil {
		fmt.Printf("Failed to parse the provided SCV file.\n")
		os.Exit(1)
	}

	problems := parseLine(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	fmt.Println(quiz(problems, timer))

}

func readFile(csvFilename *string) ([][]string, error) {
	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("failed to open CSV file: %s\n", *csvFilename)
		os.Exit(1)
	}
	defer file.Close()
	r := csv.NewReader(file)
	return r.ReadAll()
}

func parseLine(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func quiz(problems []problem, timer *time.Timer) string {
	tally := 0

problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%v: %v = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				tally++
			}
		}
	}
	return fmt.Sprintf("\nYou scored %v out of %v\n", tally, len(problems))

}
