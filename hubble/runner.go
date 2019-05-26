package hubble

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/olivere/elastic"
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

	client, err := elastic.NewClient(elastic.SetURL("http://172.16.19.24:9200"))
	if err != nil {
		log.Println(err.Error())
	}

	for {

		lastBook, err := models.
			Books(qm.OrderBy("id desc"), qm.Limit(1)).
			OneG(context.Background())

		if err != nil {
			log.Println(err.Error())
		}

		for id := 1; id <= lastBook.ID; id++ {
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
				NewBookByModel(res, client)
			}

		}
	}

}
