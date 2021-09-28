package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	lines, err := readFile(csvFilename)
	if err != nil {
		fmt.Printf("Failed to parse the provided SCV file.\n")
		os.Exit(1)
	}

	problems := parseLine(lines)
	fmt.Println(quiz(problems))

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

func quiz(problems []problem) string {
	tally := 0
	for i, p := range problems {
		fmt.Printf("Problem #%v: %v = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			tally++
		}
	}
	return fmt.Sprintf("\nYou scored %v out of %v\n", tally, len(problems))

}
