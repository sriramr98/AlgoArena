package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LogError(err error) {
	if gin.Mode() != gin.ReleaseMode {
		log.Println(err)
	}
}
