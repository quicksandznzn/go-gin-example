// Created by quicksandzn@gmail.com on 2018/12/19
package db

import (
	"../request"
	_ "github.com/go-sql-driver/mysql"
)

var db = dbInit()

type BookDto struct {
	Id     int
	Title  string
	Author string
}

func QueryById(id string) Result {
	var book BookDto
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

func Insert(book request.Book) Result {
	var result Result

	stmt, err := db.Prepare("INSERT INTO book(title,author) VALUES( ?,? )")
	checkError(err)
	defer stmt.Close()
	res, err := stmt.Exec(book.Title, book.Author)
	checkError(err)
	num, err := res.RowsAffected()
	checkError(err)
	if num > 0 {
		result = Result{
			Data:   nil,
			Status: "success",
		}
	} else {
		result = Result{
			Data:   nil,
			Status: "fail",
		}
	}

	return result
}
