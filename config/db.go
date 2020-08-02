package config

import (
	"github.com/jinzhu/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=db password=postgres")
	if err != nil {
		Logger().Error(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

