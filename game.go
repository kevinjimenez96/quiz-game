package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func play(quizItems []QuizItem) error {
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
		}
		fmt.Println()
	}

	if err := reader.Err(); err != nil {
		return err
	}

	fmt.Printf("Score %d/%d\n", correctCount, len(quizItems))
	return nil
}

func cleanInput(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}
