// Created by quicksandzn@gmail.com on 2018/12/19
package db

import (
	_ "github.com/go-sql-driver/mysql"
)

var db = dbInit()

type Book struct {
	Id     int
	Title  string
	Author string
}

func QueryById(id string) Result {
	var book Book
	var result Result
	row := db.QueryRow("select id,title,author from book where id=?", id)
	err := row.Scan(&book.Id, &book.Title, &book.Author)

	if checkError(err) {
		result = Result{
			Data:   nil,
			Status: "fail",
		}
	} else {
		result = Result{
			Data:   book,
			Status: "ok",
		}
	}

	return result
}
