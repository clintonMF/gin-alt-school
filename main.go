package main

import "github.com/gin-gonic/gin"

type Recipe struct {
	Name         string   `json:"Name"`
	Keywords     []string `json:"Keywords"`
	Ingredients  []string `json:"Ingredients"`
	Instructions string   `json:"Instructions"`
	PublishedAt  string   `json:"PublishedAt"`
}

func main() {
	router := gin.Default()
	router.Run()
}
