package item

import (
	"github.com/whs-dot-hk/go-rs3-gem-bag-calculator/price"
)

const (
	UncutSapphireCode    = 1623
	UncutEmeraldCode     = 1621
	UncutRubyCode        = 1619
	UncutDiamondCode     = 1617
	UncutDragonstoneCode = 1631
)

type Item struct {
	Code    int
	getJson func() price.Json
}

func (i *Item) getDetailJson() price.Json {
	return price.NewDetailJson(i.Code)
}

func (i *Item) getGraphJson() price.Json {
	return price.NewGraphJson(i.Code)
}

func (i *Item) GetPrice() int {
	json := i.getJson()
	return price.GetPrice(json)
}

func (i *Item) UseGetDetailJson() {
	i.getJson = func() price.Json {
		return i.getDetailJson()
	}
}

func (i *Item) UseGetGraphJson() {
	i.getJson = func() price.Json {
		return i.getGraphJson()
	}
}

func NewUncutSapphire() *Item {
	i := &Item{Code: UncutSapphireCode}
	i.UseGetDetailJson()
	return i
}

func NewUncutEmerald() *Item {
	i := &Item{Code: UncutEmeraldCode}
	i.UseGetGraphJson()
	return i
}

func NewUncutRuby() *Item {
	i := &Item{Code: UncutRubyCode}
	i.UseGetGraphJson()
	return i
}

func NewUncutDiamond() *Item {
	i := &Item{Code: UncutDiamondCode}
	i.UseGetGraphJson()
	return i
}

func NewUncutDragonstone() *Item {
	i := &Item{Code: UncutDragonstoneCode}
	i.UseGetGraphJson()
	return i
}
