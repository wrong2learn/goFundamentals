package main

import (
	"fmt"

	"github.com/wrong2learn/go-ep3-methods/internals/models"
)

func main() {
	//Value vs Pointer Receivers
	w := models.Wallet{
		Balance: 100.0,
	}
	w.Add(50)
	fmt.Println(w.Balance) // 100 (unchanged)

	w.AddPtr(50)
	fmt.Println(w.Balance) // 150

	//Method Sets
	u := models.User{
		Name: "Alice",
	}

	u.GetName()      // OK
	u.SetName("Bob") // Go auto converts to pointer

	fmt.Println(u.Name) // Bob

	//Fluent API
	//us := &models.User{}
	//u.WithName("Alice")
	//or you can write:
	us := (&models.User{}).
		WithName("Alice")

	fmt.Println(us.Name) // Alice
}
