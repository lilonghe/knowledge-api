package config

import (
	"log"

	"github.com/jinzhu/gorm"
)

type store struct {
	db *gorm.DB
}

func (this *store) Master() *gorm.DB {
	if this.db != nil {
		return this.db
	}
	db, err := gorm.Open("mysql", "root:lilonghe@tcp(127.0.0.1:3306)/knowledge?loc=Asia%2FShanghai&charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	db.DB().SetMaxOpenConns(50)
	db.DB().SetMaxIdleConns(20)
	this.db = db

	return db
}

var (
	Store store
)

func Init() {
	initDataBase()
}

func initDataBase() {
	Store := store{}
	err := Store.Master().DB().Ping()

	if err != nil {
		panic(err)
	}
}
