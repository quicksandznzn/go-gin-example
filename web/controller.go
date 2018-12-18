// Created by quicksandzn@gmail.com on 2018/12/18
package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Welcome(context *gin.Context)  {
	firstname := context.DefaultQuery("firstname", "Guest")
	lastname := context.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
