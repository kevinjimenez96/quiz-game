package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func play(quizItems []QuizItem, done chan<- string, errChan chan error) {
	reader := bufio.NewScanner(os.Stdin)
	correctCount := 0
	for _, item := range quizItems {
		fmt.Printf("%s:", item.Question)

		if !reader.Scan() {
			break
		}

		userAnswer := cleanInput(reader.Text())
		if userAnswer == item.Answer {
			correctCount++
			done <- fmt.Sprintf("Score %d/%d\n", correctCount, len(quizItems))
		}
	}

	if err := reader.Err(); err != nil {
		errChan <- err
		return
	}

	fmt.Printf("Score %d/%d\n", correctCount, len(quizItems))
	close(done)
}

func cleanInput(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}
