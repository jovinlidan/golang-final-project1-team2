package config

import (
	"fmt"
	"golang-final-project1-team2/httpserver/repositories/models"

	"github.com/jinzhu/gorm"

	_ "gorm.io/driver/mysql"
)
const (
	host           = "localhost"
	port           = "3306"
	user           = "root"
	password       = "rroooott"
	dbname         = "finalproject1"
	dbMaxIdle      = 4
	dbMaxOpenConns = 25
)


func ConnectMySQLGORM() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user,password, host, port, dbname)

	db, err := gorm.Open("mysql", dsn)

	db.AutoMigrate(&models.Todo{})
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}


	return db, nil
}
