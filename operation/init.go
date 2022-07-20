package operation

import (
	"Cocome/entity"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = entity.Db
}
