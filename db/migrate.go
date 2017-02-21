package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"music_app/models"
	)

func main(){
	db, _ := gorm.Open("mysql", "root@/music_app?charset=utf8&parseTime=True")
	db.CreateTable(&models.User{})
}
