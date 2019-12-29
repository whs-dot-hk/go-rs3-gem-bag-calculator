package price

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Json interface {
	getUrl() string
	getPrice([]byte) int
}

func GetPrice(j Json) int {
	url := j.getUrl()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return j.getPrice(body)
}
