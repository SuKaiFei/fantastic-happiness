package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// 股票
type Stock struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Date          time.Time          `json:"date" bson:"date"`                     // 股票时间，一天一个股票只能有一个
	StockExchange string             `json:"stock_exchange" bson:"stock_exchange"` // pb.StockExchange 股票所在的交易所
	StockCode     string             `json:"stock_code" bson:"stock_code"`         // 股票代码
	StockName     string             `json:"stock_name" bson:"stock_name"`         // 股票名称
	Prices        []*StockPrice      `json:"prices" bson:"prices"`                 // 股票价格
	MktCap        float64            `json:"mkt_cap" bson:"mkt_cap"`               // 总市值(万)
	Nmc           float64            `json:"nmc" bson:"nmc"`                       // 流通值(万)
	TurnoverRatio float64            `json:"turnover_ratio" bson:"turnover_ratio"` // 换手率
	PBRatio       float64            `json:"pb_ratio" bson:"pb_ratio"`             // 市净率
}

func (s *Stock) CollectionName() string {
	return "stock"
}

type StockPrice struct {
	Time  time.Time `json:"time" bson:"time"`
	Price float64   `json:"price" bson:"price"`
}
