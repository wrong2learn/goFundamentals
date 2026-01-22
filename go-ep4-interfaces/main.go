package main

import "github.com/wrong2learn/go-ep4-interfaces/internal/services"

func main() {
	email := services.EmailNotifier{Email: "alice@example.com"}
	userService := services.NewUserService(email)

	userService.RegisterUser("Alice")
}
