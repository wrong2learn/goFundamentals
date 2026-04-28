package services

import "fmt"

type SlackNotifier struct {
	Channel string
}

func (s SlackNotifier) Notify(message string) error {
	fmt.Println("Sending SLACK message to", s.Channel, ":", message)
	return nil
}
