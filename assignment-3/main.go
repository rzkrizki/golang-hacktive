package main

import (
	"golang-hacktive/server/route"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("server/controllers/views/html/templates/*.html") //static path
	app := route.NewRoute(router)
	app.SetupRouter(":7777")
}
