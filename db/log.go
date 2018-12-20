// Created by quicksandzn@gmail.com on 2018/12/20
package db

import (
	"log"
)



func checkError(err error) bool {
	if err != nil {
		log.Fatal(err.Error())
		return true
	}
	return false
}