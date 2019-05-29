package flex

import (
	"log"
	"os"
	"sync"

	"github.com/olivere/elastic"
)

var (
	esClient *elastic.Client
	once     sync.Once
)

func init() {
	once.Do(func() {
		var err error
		esClient, err = elastic.NewClient(
			elastic.SetURL("http://172.16.19.24:9200"),
			elastic.SetTraceLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		)
		if err != nil {
			log.Println(err.Error())
		}
	})
}
