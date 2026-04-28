package services

import "fmt"

type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) Notify(message string) error {
	fmt.Println("Sending EMAIL to", e.Email, ":", message)
	return nil
}
