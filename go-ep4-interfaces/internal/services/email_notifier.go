package services

import "fmt"

type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) Notify(message string) error {
	fmt.Print("Sending email to: ", e.Email, " : ", message)
	return nil
}
