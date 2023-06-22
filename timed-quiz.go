package main

import (
	"bufio" // Buffer I/O
	"flag"  // Runtime Flags
	"fmt"   // for printing
	"log"   //for logging errors
	"os"    // Reading in raw bytes
	"strings"
)

// Return Quiz PTR to practice memory in GO
func readCSV(path string) (*Quiz, error) {
	file, err := os.Open(path)
	if err != nil {
		//log error
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	quiz := &Quiz{
		QuestionCount: 0,
		Correct:       0,
		Incorrect:     0,
		Questions:     []*Question{},
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//create copy of current line
		line := scanner.Text()

		//get csv items from line, don't split on commas in quotes
		items := strings.Split(line, ",")
		//create question object
		question := &Question{
			Question: items[0],
			Answer:   items[1],
		}
		//add question to quiz
		quiz.Questions = append(quiz.Questions, question)
		quiz.QuestionCount++
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file: %w", err)
	}
	return quiz, nil
}

func main() {
	//Define quiz flag -- flag.String("name of flag", "default value", "Flag Usage/Description") stored as ptr - need to derefernce to get value.
	var quizPathPtr = flag.String("quiz", "problems.csv", "name of CSV file to read in for quiz")

	//Parse flag(s)
	flag.Parse()

	//read in csv and generate quiz object
	quiz, err := readCSV(*quizPathPtr)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
		return
	}

	println("Press enter to start quiz, you have 30 seconds to complete")
	//wait for user to press enter
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	for i := 0; i < quiz.QuestionCount; i++ {
		fmt.Println("Question: " + quiz.Questions[i].Question)
		//wait for user to enter answer
		answer, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		//remove newline from answer
		answer = strings.TrimSuffix(answer, "\n")
		//check if answer is correct
		if answer == quiz.Questions[i].Answer {
			quiz.Correct++
		} else {
			quiz.Incorrect++
		}
	}
	quiz.PrintQuizResults()
}

type Question struct {
	Question string
	Answer   string
}

type Quiz struct {
	QuestionCount int
	Correct       int
	Incorrect     int
	Questions     []*Question
}

func (q Quiz) PrintQuizResults() {
	red := "\033[31m"
	green := "\033[32m"
	reset := "\033[0m"

	fmt.Printf(red+"Incorrect Questions: %v/%v\n", q.Incorrect, q.QuestionCount)
	fmt.Printf(green+"Correct Questions: %v/%v\n"+reset, q.Correct, q.QuestionCount)
}
