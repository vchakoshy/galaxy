package rest

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PastVersion(c *gin.Context) {

	u := "https://api2.fidibo.com" + c.Request.URL.Path

	req, err := http.NewRequest("POST", u, c.Request.Body)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)

}
