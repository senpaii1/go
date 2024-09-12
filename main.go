package main

import (
	u "main/abhinav"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/",func (c *gin.Context){
		u.GinFunction(c)
	})
	router.Run(":5000")
	router.Use()
}
