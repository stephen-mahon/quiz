package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var title = "Exercise #1: Quiz Game"
var help = "Load file name with -f. See README for more information."
var errCLI = "You must enter arguments! Type -help for help."

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(errCLI)
	} else if len(args) == 1 && args[0] == "-help" {
		fmt.Println(title)
		fmt.Println(help)
	} else if len(args) == 2 && args[0] == "-f" {
		dat := readfile(args[1])
		sols, ans := quiz(dat)

		output, err := checkSol(sols, ans)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(output)
		}
	} else {
		fmt.Println(errCLI)
	}
}

func readfile(fileName string) []string {
	file, err := os.Open(fileName + ".csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var output []string

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}

func quiz(dat []string) ([]string, []string) {
	var input, solutions []string
	for i := range dat {
		question := strings.Split(dat[i], ",")[0]
		ans := strings.Split(dat[i], ",")[1]
		fmt.Printf("%v:\t", question)
		input = append(input, readInput())
		solutions = append(solutions, ans)
	}

	return solutions, input
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func checkSol(sol, ans []string) (string, error) {
	if len(sol) != len(ans) {
		return "", errors.New("solution length and answer input length do not match")
	}
	var tally int
	for i := range sol {
		if sol[i] == ans[i] {
			tally++
		}
	}
	output := fmt.Sprintf("\nYou scored %v out of %v\n", tally, len(sol))
	return output, nil
}
