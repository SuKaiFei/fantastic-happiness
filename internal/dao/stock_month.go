package dao

import (
	"context"
	"github.com/SuKaiFei/fantastic-happiness/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type StockMonther interface {
	AddStockMonth(context.Context, *model.Stock) (*mongo.UpdateResult, error)
}

func (d *dao) AddStockMonth(ctx context.Context, stock *model.Stock) (result *mongo.UpdateResult, err error) {
	return nil, err
}
