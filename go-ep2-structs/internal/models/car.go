package models

type Car struct {
	Engine // embedded struct
	Brand  string
}
