package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// 股票
type StockMin struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Date          time.Time          `json:"date" bson:"date"`                     // 股票时间，一天一个股票只能有一个
	StockExchange string             `json:"stock_exchange" bson:"stock_exchange"` // 股票所在的交易所
	StockCode     string             `json:"stock_code" bson:"stock_code"`         // 股票代码
	StockName     string             `json:"stock_name" bson:"stock_name"`         // 股票名称
	Prices        []*StockPriceItem  `json:"prices" bson:"prices"`                 // 股票价格
}

func (s *StockMin) CollectionName() string {
	return "stock_min"
}

type StockPriceItem struct {
	Time   time.Time `json:"time" bson:"time"`
	High   float64   `json:"high" bson:"high"`
	Volume int       `json:"volume" bson:"volume"`
	Open   float64   `json:"open" bson:"open"`
	Close  float64   `json:"close" bson:"close"`
	Low    float64   `json:"low" bson:"low"`
}
