package domains

import (
	"chiku3_app/domains"
//	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type StaffRepositoryMysql struct {
	gorm.Model
	UserId   int
	UserName string
}

// func (s StaffRepositoryMysql) InitDataBase() {
// 	db, err := gorm.Open(mysql.Open())
// }
