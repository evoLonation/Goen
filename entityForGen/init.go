package entityForGen

import (
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

//func init() {
//	dsn := "root:2002116yy@tcp(127.0.0.1:3306)/Cocome?charset=utf8mb4&parseTime=True&loc=Local"
//	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	err := db.AutoMigrate(&Item{}, &CashDesk{})
//	if err != nil {
//		print(err)
//		return
//	}
//	Db = db
//}
