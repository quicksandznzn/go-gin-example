// Created by quicksandzn@gmail.com on 2018/7/25
package main

import (
	"github.com/gin-gonic/gin"
	"./web"
)


func main() {
	router := gin.Default()
	router.GET("/welcome", web.Welcome)
	router.POST("/insert", web.Insert)
	router.GET("/getById", web.GetById)
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


