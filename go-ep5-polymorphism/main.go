package main

import "github.com/wrong2learn/go-ep5-polymorphism/internal/services"

func main() {

	email := services.EmailNotifier{
		Email: "admin@example.com",
	}

	sms := services.SMSNotifier{
		Phone: "+44123456789",
	}

	slack := services.SlackNotifier{
		Channel: "#general",
	}

	userService1 := services.NewUserService(email)
	userService1.RegisterUser("Alice")

	userService2 := services.NewUserService(sms)
	userService2.RegisterUser("Bob")

	userService3 := services.NewUserService(slack)
	userService3.RegisterUser("Charlie")
}
