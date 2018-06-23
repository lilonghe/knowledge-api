package config

import (
	"log"

	"github.com/jinzhu/gorm"
)

type store struct {
}

func (this *store) Master() *gorm.DB {
	db, err := gorm.Open("mysql", "root:lilonghe@tcp(127.0.0.1:3306)/knowledge?loc=Asia%2FShanghai&charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)
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
	Store.Master()
	//db := Store.Master()
}
