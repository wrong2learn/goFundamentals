package services

import "fmt"

type SMSNotifier struct {
	Phone string
}

func (s SMSNotifier) Notify(message string) error {
	fmt.Println("Sending SMS to", s.Phone, ":", message)
	return nil
}
