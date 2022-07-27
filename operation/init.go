package operation

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {

	var err error
	db, err = sqlx.Open("mysql", "root:2002116yy@tcp(127.0.0.1:3306)/Cocome")
	if err != nil {
		print(err)
	}
}
