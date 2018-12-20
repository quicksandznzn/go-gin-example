// Created by quicksandzn@gmail.com on 2018/12/19
package db

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db = dbInit()

type Book struct {
	Id              sql.NullInt64
	Title          sql.NullString
	Author            sql.NullString

}


func QueryById(id string) gin.H{
	var book Book
	var result gin.H

	row := db.QueryRow("select id,title,author from book where id=?", id)
	err :=row.Scan(&book.Id,&book.Title,&book.Author)

	if checkError(err) {
		result = gin.H{
			"data":   nil,
			"status": "failed",
		}
	} else {
		result = gin.H{
			"data":   book,
			"status": "ok",
		}
	}

	return result
}
