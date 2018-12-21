// Created by quicksandzn@gmail.com on 2018/12/21
package request

type Book struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
