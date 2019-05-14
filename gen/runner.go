package gen

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

// Run runs generator
func Run() {
	var db *sqlx.DB
	var err error

	mysqlDNS := "reader:j3AhwCj4SqVQI62Y@tcp(79.175.173.69:3306)/fidibo1_fidibo?autocommit=true"
	db, err = sqlx.Open("mysql", mysqlDNS)
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
		return
	}

	r := db.QueryRowx(`select * from fidibo1_fidibo.author;`)
	col, err := r.Columns()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(col)

	coltype, err := r.ColumnTypes()
	if err != nil {
		log.Println(err.Error())
		return
	}
	for _, ct := range coltype {
		// lp(ct.DatabaseTypeName())
		// lp(ct.Length())
		lp(ct.Name())
		// lp(ct.ScanType())
	}
}

func lp(o ...interface{}) {
	log.Println(o)
}
