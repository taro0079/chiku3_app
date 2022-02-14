package main

import (
	// "chiku3_app/domains"
	"chiku3_app/domains"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	UserId   int
	UserName string
}

func main() {
	// test, err := domains.NewUserName("ch")
	// if err != nil {
	// 	panic("error")
	// }
	//
	//
	// fmt.Println(test)
	repo := domains.StaffRepositoryMysql
	db := connectDb()
	db.AutoMigrate(&User{})
}

func connectDb() *gorm.DB {
	connection := "user1:takano0820@tcp(127.0.0.1:3306)/go_sample"
	db, err := gorm.Open(mysql.Open(connection))
	if err != nil {
		fmt.Println("error during connection")
	}
	return db

}
