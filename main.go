package main

import (
	"github.com/gin-gonic/gin"

	"github.com/oyachi/goscraper/controller"
)

func main() {
	router := gin.Default()

	router.GET("/weather", controller.GetWeatherInfo)

	router.Run()
}
