package controllers

import (
	"golang-hacktive/server/controllers/views"

	"github.com/gin-gonic/gin"
)

func WriteJsonResponse(ctx *gin.Context, response *views.Response) {
	ctx.JSON(response.Status, response)
}
