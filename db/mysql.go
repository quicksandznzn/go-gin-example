// Created by quicksandzn@gmail.com on 2018/12/19
package db

import (
	"database/sql"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Mysql struct {
		Path       string `yaml:"path"`
		DriverName string `yaml:"driverName"`
	}
}

const ConfigPath = "./config/conf.yml"

func dbInit() *sql.DB {

	config := Config{}

	buffer, err := ioutil.ReadFile(ConfigPath)
	checkError(err)

	err = yaml.Unmarshal(buffer, &config)
	checkError(err)

	db, err := sql.Open(config.Mysql.DriverName, config.Mysql.Path)
	checkError(err)

	err = db.Ping()
	checkError(err)
	log.Println("DB connect successfully ============")

	return db
}
