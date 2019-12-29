package price

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

const graphJsonUrl = "http://services.runescape.com/m=itemdb_rs/api/graph/%d.json"

type GraphJson struct {
	ItemCode int
}

func NewGraphJson(itemCode int) *GraphJson {
	return &GraphJson{ItemCode: itemCode}
}

func (j *GraphJson) getUrl() string {
	return fmt.Sprintf(graphJsonUrl, j.ItemCode)
}

type CurrentObj struct {
	Price int
}

type GraphObj struct {
	Daily map[string]int
}

func (o *GraphObj) GetPrice() int {
	var largestTimestamp int
	for k, _ := range o.Daily {
		i, err := strconv.Atoi(k)
		if err != nil {
			log.Fatal(err)
		}
		if i > largestTimestamp {
			largestTimestamp = i
		}
	}
	largestTimestampKey := fmt.Sprintf("%d", largestTimestamp)
	return o.Daily[largestTimestampKey]
}

func (j *GraphJson) getPrice(body []byte) int {
	var graph GraphObj
	err := json.Unmarshal(body, &graph)
	if err != nil {
		log.Fatal(err)
	}
	return graph.GetPrice()
}
