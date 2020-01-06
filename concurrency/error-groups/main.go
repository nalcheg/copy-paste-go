package main

import (
	"errors"
	"log"
	"time"
)

func main() {
}

type App struct {
	first  string
	second string
	third  string
}

var errorFirst, errorSecond, errorThird error

func init() {
	errorFirst = errors.New("just an error | first")
	errorSecond = errors.New("just an error | second")
	errorThird = errors.New("just an error | third")
}

func NewApp(first string, second string, third string) *App {
	return &App{first: first, second: second, third: third}
}

func (a *App) GetFirst() (string, error) {
	time.Sleep(300 * time.Millisecond)
	log.Print("first")

	var reason int
	reason = 1
	if reason == 1 {
		return a.first, nil
	} else {
		return "", errorFirst
	}
}

func (a *App) GetSecond() (string, error) {
	time.Sleep(200 * time.Millisecond)
	log.Print("second")

	var reason int
	reason = 0
	if reason == 1 {
		return a.second, nil
	} else {
		return "", errorSecond
	}
}

func (a *App) GetThird() (string, error) {
	time.Sleep(100 * time.Millisecond)
	log.Print("third")

	var reason int
	reason = 1
	if reason == 1 {
		return a.third, nil
	} else {
		return "", errorThird
	}
}
