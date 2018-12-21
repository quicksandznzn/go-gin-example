// Created by quicksandzn@gmail.com on 2018/7/25
package main

import (
	"./web"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.Use(Auth())
	{
		v1.GET("/welcome", web.Welcome)
		v1.POST("/insert", web.Insert)
		v1.GET("/getById", web.GetById)
	}

	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":8881")
	// router.Run(":3000") for a hard coded port

}
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		//Get token and e-mail from header
		token := context.Request.Header.Get("token")

		//check to see if email & token were provided
		if len(token) == 0 {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		} else {
			context.Next()
		}
	}
}
