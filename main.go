package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type QuizItem struct {
	question string
	answer   string
}

func main() {
	questionsFileFlag := flag.String("questions", "problems.csv", "custom question's file")
	quizItems, err := loadQuizItems(*questionsFileFlag)
	if err != nil {
		log.Fatal(err)
	}

	err = play(quizItems)
	if err != nil {
		log.Fatal(err)
	}
}

func play(quizItems []QuizItem) error {
	reader := bufio.NewScanner(os.Stdin)
	correctCount := 0
	for _, item := range quizItems {
		fmt.Printf("%s:", item.question)
		reader.Scan()
		userAnswer := cleanInput(reader.Text())
		if userAnswer == item.answer {
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
	return strings.TrimSpace(input)
}

func loadQuizItems(filename string) ([]QuizItem, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var quizItems []QuizItem

	for {
		rawItem, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		quizItems = append(quizItems, QuizItem{
			question: rawItem[0],
			answer:   rawItem[1],
		})

	}

	return quizItems, nil
}
