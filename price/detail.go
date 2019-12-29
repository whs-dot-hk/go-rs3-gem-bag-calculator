package price

import (
	"encoding/json"
	"fmt"
	"log"
)

const detailJsonUrl = "http://services.runescape.com/m=itemdb_rs/api/catalogue/detail.json?item="

type DetailJson struct {
	ItemCode int
}

func NewDetailJson(itemCode int) *DetailJson {
	return &DetailJson{ItemCode: itemCode}
}

func (j *DetailJson) getUrl() string {
	return fmt.Sprintf("%s%d", detailJsonUrl, j.ItemCode)
}

type ItemObj struct {
	Current CurrentObj
}

type DetailObj struct {
	Item ItemObj
}

func (o *DetailObj) GetPrice() int {
	return o.Item.Current.Price
}

func (j *DetailJson) getPrice(body []byte) int {
	var detail DetailObj
	err := json.Unmarshal(body, &detail)
	if err != nil {
		log.Fatal(err)
	}
	return detail.GetPrice()
}
