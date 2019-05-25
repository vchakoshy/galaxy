package hubble

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)

func Run() {

	mysqlDNS := "reader:j3AhwCj4SqVQI62Y@tcp(79.175.173.69:3306)/fidibo1_fidibo?autocommit=true&parseTime=true"
	db, err := sqlx.Open("mysql", mysqlDNS)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()

	boil.SetDB(db)
	// boil.DebugMode = true

	id := 1
	for {
		res, err := models.
			Books(
				qm.Load("BookStat"),
				qm.Load("Publisher"),
				qm.Load("BookCategoryAssigns"),
				qm.Load("Author"),
				qm.Where("id=?", id)).
			OneG(context.Background())

		if err != nil {
			log.Println(err.Error())
		}

		if err == nil {
			NewBookByModel(res)
		}

		id++
	}

	fmt.Println("elastic called")
}
