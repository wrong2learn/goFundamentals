package services

import "fmt"

type Logger struct{}

func (l Logger) Log(msg string) {
	fmt.Println("[LOG]", msg)
}
