package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWeatherInfo(c *gin.Context) {
	m, err := Scrape()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(http.StatusOK, string(m))
}