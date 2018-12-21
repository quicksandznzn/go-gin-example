// Created by quicksandzn@gmail.com on 2018/12/18
package web

import (
	"../db"
	"../request"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func Welcome(context *gin.Context) {
	context.String(http.StatusOK, "welcome")
}

func Insert(context *gin.Context) {
	var book request.Book
	if context.BindWith(&book, binding.JSON) == nil {
		context.JSON(http.StatusOK, db.Insert(book))
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "binding json error"})
	}

}

func GetById(context *gin.Context) {
	id := context.Query("id")
	context.JSON(http.StatusOK, db.QueryById(id))
}
