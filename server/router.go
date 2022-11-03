package server

import (
	"akaflieg/fresskasse/v2/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Static("/img", "./assets/img")
	router.StaticFile("/favicon.ico", "./assets/img/favicon.ico")

	router.LoadHTMLGlob("templates/**/*")

	home := new(controllers.HomeController)
	router.GET("/", home.Retrieve)

	all := new(controllers.AllController)
	router.GET("/all", all.Retrieve)

	notFound := new(controllers.NotFoundController)
	router.NoRoute(notFound.Retrieve)

	return router
}
