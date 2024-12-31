package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Quiz configuration Constant
const (
	TIME_PER_QUESTION = 30 * time.Second // 30 seconds per question
	EXIT_COMMAND      = "exit"
	PASS_PERCENTAGE   = 60.0
)

// Performance levels constant
const (
	EXCELLENT = 90.0
	GOOD      = 75.0
	AVERAGE   = 60.0
)

// Question dataType
type Question struct {
	Text          string
	Options       []string
	CorrectAnswer int
}

// Quiz to manage examination system
type Quiz struct {
	Questions      []Question
	Score          int
	TotalQuestions int
	Scanner        *bufio.Scanner
}

func NewQuiz() *Quiz {
	questions := []Question{
		{
			Text: "What is the capital of India?",
			Options: []string{
				"Mumbai",
				"New Delhi",
				"Bangalore",
				"Kolkata",
			},
			CorrectAnswer: 2,
		},
		{
			Text: "Which planet is known as the Red Planet?",
			Options: []string{
				"Venus",
				"Jupiter",
				"Mars",
				"Saturn",
			},
			CorrectAnswer: 3,
		},
		{
			Text: "What is 2 + 2 × 3?",
			Options: []string{
				"12",
				"8",
				"10",
				"8",
			},
			CorrectAnswer: 2,
		},
		{
			Text: "Who wrote 'Discovery of India'?",
			Options: []string{
				"Mahatma Gandhi",
				"Jawaharlal Nehru",
				"Rabindranath Tagore",
				"Sardar Patel",
			},
			CorrectAnswer: 2,
		},
		{
			Text: "Which is the largest ocean in the world?",
			Options: []string{
				"Indian Ocean",
				"Atlantic Ocean",
				"Arctic Ocean",
				"Pacific Ocean",
			},
			CorrectAnswer: 4,
		},
	}

	return &Quiz{
		Questions:      questions,
		TotalQuestions: len(questions),
		Scanner:        bufio.NewScanner(os.Stdin),
	}
}

func (q *Quiz) displayQuestion(questionNum int, question Question) {
	fmt.Printf("\nQuestion %d/%d:\n", questionNum+1, q.TotalQuestions)
	fmt.Println(question.Text)
	for i, option := range question.Options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	fmt.Printf("\nEnter your answer (1-%d) or '%s' to quit: ", len(question.Options), EXIT_COMMAND)
}

func (q *Quiz) readAnswer() (int, error) {
	q.Scanner.Scan()
	input := strings.TrimSpace(q.Scanner.Text())

	if strings.ToLower(input) == EXIT_COMMAND {
		return -1, nil
	}

	answer, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("please enter a valid number")
	}

	return answer, nil
}

func getPerformanceLevel(percentage float64) string {
	switch {
	case percentage >= EXCELLENT:
		return "Excellent"
	case percentage >= GOOD:
		return "Good"
	case percentage >= AVERAGE:
		return "Average"
	default:
		return "Needs Improvement"
	}
}

func (q *Quiz) StartQuiz() {
	fmt.Println("\nWelcome to the Online Examination System!")
	fmt.Printf("You have %v per question. Total questions: %d\n", TIME_PER_QUESTION, q.TotalQuestions)
	fmt.Println("Press Enter to start the quiz...")
	q.Scanner.Scan()

	for i, question := range q.Questions {
		timer := time.NewTimer(TIME_PER_QUESTION)
		answerChan := make(chan int)
		timeoutChan := make(chan bool)

		q.displayQuestion(i, question)

		go func() {
			for {
				answer, err := q.readAnswer()
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					continue
				}
				answerChan <- answer
				break
			}
		}()

		go func() {
			<-timer.C
			timeoutChan <- true
		}()

		select {
		case answer := <-answerChan:
			timer.Stop()
			if answer == -1 {
				fmt.Println("\nQuiz terminated by user.")
				q.showResults()
				return
			}
			if answer < 1 || answer > len(question.Options) {
				fmt.Println("Invalid option selected. No points awarded.")
				continue
			}
			if answer == question.CorrectAnswer {
				fmt.Println("Correct!")
				q.Score++
			} else {
				fmt.Printf("Incorrect. The correct answer was: %d\n", question.CorrectAnswer)
			}

		case <-timeoutChan:
			fmt.Println("\nTime's up! Moving to next question...")
		}
	}

	q.showResults()
}

func (q *Quiz) showResults() {
	percentage := (float64(q.Score) / float64(q.TotalQuestions)) * 100
	performance := getPerformanceLevel(percentage)

	fmt.Println("\n--- Quiz Results ---")
	fmt.Printf("Total Questions: %d\n", q.TotalQuestions)
	fmt.Printf("Correct Answers: %d\n", q.Score)
	fmt.Printf("Percentage: %.2f%%\n", percentage)
	fmt.Printf("Performance: %s\n", performance)

	if percentage >= PASS_PERCENTAGE {
		fmt.Println("Congratulations! You passed the quiz!")
	} else {
		fmt.Println("Keep practicing. You can do better!")
	}
}

func main() {
	quiz := NewQuiz()
	quiz.StartQuiz()
}
