package operation

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {

	var err error
	Db, err = sqlx.Open("mysql", "root:2002116yy@tcp(127.0.0.1:3306)/entity_manager")
	if err != nil {
		print(err)
	}
}
