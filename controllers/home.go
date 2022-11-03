package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (h HomeController) Retrieve(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home.html",
		gin.H{
			"title": "Startseite",
		},
	)
}
