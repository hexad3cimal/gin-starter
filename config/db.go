package config

import (
	"github.com/jinzhu/gorm"

)

var DB *gorm.DB
var log = InitLogger()
func InitDB(){
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=db password=postgres")
	if err == nil {
		db.DB().SetMaxIdleConns(10)
		db.LogMode(true)
		DB = db
	}
	log.Error("Db failure")

}

func GetDB() *gorm.DB {
	return DB
}

