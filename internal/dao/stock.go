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

type Stocker interface {
	AddStock(context.Context, *model.Stock) (*mongo.UpdateResult, error)
	GetStockByToday(context.Context, ...string) ([]*model.Stock, error)
}

func (d *dao) GetStockByToday(ctx context.Context, projections ...string) (stocks []*model.Stock, err error) {
	var (
		curDate = time.Now()
		date    = time.Date(curDate.Year(), curDate.Month(), curDate.Day(), 0, 0, 0, 0, time.Local)
		filter  = bson.D{
			{"date", date},
		}
		opt = new(options.FindOptions).SetLimit(1)
	)

	if len(projections) > 0 {
		projection := bson.M{}
		for _, s := range projections {
			projection[s] = 1
		}
		opt.SetProjection(projection)
	}

	cursor, err := d.fHDB.Collection(new(model.Stock).CollectionName()).Find(ctx, filter, opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if err := cursor.All(ctx, &stocks); err != nil {
		return nil, errors.WithStack(err)
	}

	return
}

func (d *dao) AddStock(ctx context.Context, stock *model.Stock) (result *mongo.UpdateResult, err error) {
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
		update = bson.D{
			{"$addToSet",
				bson.M{
					"prices": bson.M{"$each": stock.Prices},
				},
			},
			{"$set",
				bson.D{
					{"stock_name", stock.StockName},
					{"mkt_cap", stock.MktCap},
					{"nmc", stock.Nmc},
					{"turnover_ratio", stock.TurnoverRatio},
					{"pb_ratio", stock.PBRatio},
				},
			},
		}
		opts = new(options.UpdateOptions).SetUpsert(true)
	)

	result, err = d.fHDB.Collection(stock.CollectionName()).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return
}
