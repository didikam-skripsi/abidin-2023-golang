package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (con HomeController) Index(c *gin.Context) {
	data := gin.H{
		"title": "hello world",
	}
	c.JSON(http.StatusOK, data)
}
