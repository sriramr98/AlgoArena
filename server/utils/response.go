package utils

import "github.com/gin-gonic/gin"

func SuccessResponse(data any) gin.H {
	return gin.H{
		"success": true,
		"data":    data,
	}
}

func FailureResponse(message string) gin.H {
	return gin.H{
		"success": false,
		"message": message,
	}
}
