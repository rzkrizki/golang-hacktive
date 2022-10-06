package controllers

import (
	"golang-hacktive/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexAlarm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", "")
}

func GenerateData(ctx *gin.Context) {

	response := services.GenerateData()
	WriteJsonResponse(ctx, response)
}

func ReadData(ctx *gin.Context) {
	response := services.ReadData()
	WriteJsonResponse(ctx, response)
}
