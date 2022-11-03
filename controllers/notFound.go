package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotFoundController struct{}

func (n NotFoundController) Retrieve(c *gin.Context) {
	c.HTML(
		http.StatusNotFound,
		"404.html",
		gin.H{
			"title": "Seite nicht gefunden",
		},
	)
}
