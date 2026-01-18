package main

import (
	"encoding/json"
	"fmt"

	"github.com/wrong2learn/go-ep2-structs/internal/models"
	"github.com/wrong2learn/go-ep2-structs/internal/services"
)

func main() {
	fmt.Println("Hello from main.go!")

	//Struct embedding
	car := models.Car{
		Engine: models.Engine{
			HorsePower: 200,
		},
		Brand: "Tesla",
	}

	car.Start()

	//Anonymous structs
	user := struct {
		Name string
		Age  int
	}{
		Name: "Alice",
		Age:  30,
	}
	fmt.Println(user.Name, user.Age)

	//Struct Tags
	u := models.User{
		ID:    1,
		Name:  "Alice",
		Email: "",
	}

	data, _ := json.Marshal(u)
	fmt.Println(string(data))

	//Value vs Pointer Semantics
	c := models.Counter{Value: 1}

	services.Increment(c)
	fmt.Println(c.Value) // 1 (unchanged)

	services.IncrementPtr(&c)
	fmt.Println(c.Value) // 2 (updated)

	//Logger Service
	s := services.AuthService{Name: "Auth"}
	s.Log("Service started") //print: [LOG] Service started
}
