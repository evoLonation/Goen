package entityManager

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var Db *sqlx.DB
var ItemManager *ManagerGeneric[Item, *Item]

func init() {
	var err error
	Db, err = sqlx.Open("mysql", "root:2002116yy@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	ItemManager, err = NewManager[Item, *Item]("item")
	if err != nil {
		log.Fatal(err)
	}
}
