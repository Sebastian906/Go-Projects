package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// First, connect without specifying a database to create it if it doesn't exist
	d, err := gorm.Open("mysql", "root:sebas1912@tcp(localhost:3306)/?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}

	// Create database if it doesn't exist
	d.Exec("CREATE DATABASE IF NOT EXISTS goproject")
	d.Close()

	// Now connect to the specific database
	d, err = gorm.Open("mysql", "root:sebas1912@tcp(localhost:3306)/goproject?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
