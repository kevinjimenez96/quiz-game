package main

import (
	"encoding/csv"
	"io"
	"os"
)

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

		if len(rawItem) != 2 {
			continue
		}

		quizItems = append(quizItems, QuizItem{
			Question: rawItem[0],
			Answer:   rawItem[1],
		})

	}

	return quizItems, nil
}
