package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/imshakthi/goland/models"
)

const (
	Message = "message"
)

func SendMessageWithStatus(ctx *gin.Context, httpStatusCode int, message string) {
	responseMap := models.GenericMsgResponse{
		Message: message,
	}
	ctx.JSON(httpStatusCode,
		responseMap,
	)
}
