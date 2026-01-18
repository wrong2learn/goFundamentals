package services

import "github.com/wrong2learn/go-ep2-structs/internal/models"

func Increment(c models.Counter) {
	c.Value++
}

func IncrementPtr(c *models.Counter) {
	c.Value++
}
