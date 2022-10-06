package route

import (
	"golang-hacktive/server/controllers"

	"github.com/gin-gonic/gin"
)

type router struct {
	router *gin.Engine
}

func NewRoute(r *gin.Engine) *router {
	return &router{
		router: r,
	}
}

func (r *router) SetupRouter(port string) {
	r.router.GET("/generate-data", controllers.GenerateData)
	r.router.GET("/read-data", controllers.ReadData)
	r.router.GET("/", controllers.IndexAlarm)
	r.router.Run(port)
}
