package services

import "fmt"

type DiscordNotifier struct {
	Server string
}

func (d DiscordNotifier) Notify(message string) error {
	fmt.Println("Sending DISCORD message to", d.Server, ":", message)
	return nil
}
