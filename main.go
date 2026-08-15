package main

import (
	"flag"
	"log"
)

type QuizItem struct {
	Question string
	Answer   string
}

func main() {
	questionsFileFlag := flag.String("questions", "problems.csv", "custom question's file")
	flag.Parse()

	quizItems, err := loadQuizItems(*questionsFileFlag)
	if err != nil {
		log.Fatal(err)
	}

	err = play(quizItems)
	if err != nil {
		log.Fatal(err)
	}
}
