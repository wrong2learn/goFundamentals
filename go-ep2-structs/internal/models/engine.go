package models

import "fmt"

type Engine struct {
	HorsePower int
}

func (e Engine) Start() {
	fmt.Println("Engine started with", e.HorsePower, "HP")
}
