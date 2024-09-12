package abhinav

import (
	"github.com/gin-gonic/gin"
)

func GinFunction (c *gin.Context){
	c.JSON(200, gin.H {
		"message" : "Hello Server",
	})}