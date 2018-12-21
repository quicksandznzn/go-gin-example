// Created by quicksandzn@gmail.com on 2018/12/18
package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-gonic/gin/binding"
	"../db"
)

type Book struct {
	Name string `json:"name" binding:"required"`
}

func Welcome(context *gin.Context) {
	context.String(http.StatusOK, "welcome")
}

func Insert(context *gin.Context) {
	var book Book
	if context.BindWith(&book, binding.JSON) == nil {
		context.JSON(http.StatusOK, gin.H{"JSON=== status": "you are logged in"})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "binding json error"})
	}

}

func GetById(context *gin.Context) {
	id := context.Query("id")
	context.JSON(http.StatusOK,db.QueryById(id))
}
