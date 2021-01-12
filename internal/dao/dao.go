package dao

import (
	"context"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDB, NewRedis)

// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)

	Stocker
	StockMiner
	StockDayer
	StockMonther
}

// dao dao.
type dao struct {
	db   *mongo.Client
	fHDB *mongo.Database

	redis *redis.Redis
	cache *fanout.Fanout
}

// New new a dao and return.
func New(r *redis.Redis, db *mongo.Client) (d Dao, cf func(), err error) {
	return newDao(r, db)
}

func newDao(r *redis.Redis, db *mongo.Client) (d *dao, cf func(), err error) {
	var cfg struct {
		DemoExpire xtime.Duration
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = &dao{
		db:    db,
		fHDB:  db.Database("fantastic-happiness"),
		redis: r,
		cache: fanout.New("cache"),
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	d.cache.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
