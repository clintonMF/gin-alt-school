package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	Name         string    `json:"Name"`
	Keywords     []string  `json:"Keywords"`
	Ingredients  []string  `json:"Ingredients"`
	Instructions string    `json:"Instructions"`
	PublishedAt  time.Time `json:"PublishedAt"`
}

func main() {
	router := gin.Default()
	router.Run()
}
