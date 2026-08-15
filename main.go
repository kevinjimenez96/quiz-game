package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"time"
)

type QuizItem struct {
	Question string
	Answer   string
}

func main() {
	questionsFileFlag := flag.String("questions", "problems.csv", "custom question's file")
	timerFlag := flag.Int("timer", 30, "question's timer")
	flag.Parse()

	if *timerFlag <= 0 {
		log.Fatal(errors.New("invalid timer"))
	}

	quizItems, err := loadQuizItems(*questionsFileFlag)
	if err != nil {
		log.Fatal(err)
	}

	gameDone := make(chan string)
	gameErrors := make(chan error)

	go play(quizItems, gameDone, gameErrors)

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Time exceded!")
	case <-gameDone:
	case err = <-gameErrors:
		log.Fatal(err)
	}
}
