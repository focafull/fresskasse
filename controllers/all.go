package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AllController struct{}

func (a AllController) Retrieve(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"all.html",
		gin.H{
			"title": "Alle",
		},
	)
}
