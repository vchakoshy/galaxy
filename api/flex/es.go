package flex 

import (
	"sync"
	"log"
	"github.com/olivere/elastic"
)

var (
	esClient *elastic.Client
	once sync.Once 
)

func init(){
	once.Do(func(){
		var err error 
		esClient , err = elastic.NewClient(elastic.SetURL("http://172.16.19.24:9200"))
		if err != nil {
			log.Println(err.Error())
		}
	})
}