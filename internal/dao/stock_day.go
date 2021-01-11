package dao

import (
	"context"
	"github.com/SuKaiFei/fantastic-happiness/internal/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type StockDayer interface {
	AddStockDay(context.Context, *model.StockDay) (*mongo.UpdateResult, error)
}

func (d *dao) AddStockDay(ctx context.Context, stock *model.StockDay) (result *mongo.UpdateResult, err error) {
	if stock == nil {
		return
	}

	var (
		date   = time.Date(stock.Date.Year(), stock.Date.Month(), stock.Date.Day(), 0, 0, 0, 0, time.Local)
		filter = bson.D{
			{"date", date},
			{"stock_exchange", stock.StockExchange},
			{"stock_code", stock.StockCode},
		}
		update = bson.D{}
		opts   = new(options.UpdateOptions).SetUpsert(true)
	)

	if len(stock.Prices) > 0 {
		update = append(update,
			bson.E{
				Key: "$addToSet",
				Value: bson.M{
					"prices": bson.M{"$each": stock.Prices},
				},
			})
	}
	update = append(update, bson.E{Key: "$set",
		Value: bson.D{
			{"stock_name", stock.StockName},
		},
	})

	result, err = d.fHDB.Collection(stock.CollectionName()).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return
}
