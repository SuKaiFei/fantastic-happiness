package service

import (
	"context"
	"time"

	pb "github.com/SuKaiFei/fantastic-happiness/api"
	"github.com/SuKaiFei/fantastic-happiness/internal/model"
	"github.com/go-kratos/kratos/pkg/log"

	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"go.uber.org/atomic"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	stockCollectMinTaskFlag = atomic.NewBool(true)
)

func (s *Service) SyncStockMinAtToday(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	if !stockCollectMinTaskFlag.Load() {
		return &emptypb.Empty{}, nil
	}
	stockCollectMinTaskFlag.Store(false)

	go s.SyncStockMinAtTodayJob()

	return &emptypb.Empty{}, nil
}

func (s *Service) SyncStockMinAtTodayJob() {
	defer stockCollectMinTaskFlag.Store(true)

	log.Info("SyncStockMinAtTodayJob start")
	ctx := context.Background()

	stocks, err := s.dao.GetStockByToday(ctx, "_id", "stock_exchange", "stock_name", "stock_code")
	if err != nil {
		log.Error("[SyncStockMinAtTodayJob]GetStockByToday err(%+v)", err)
		return
	}

	curTime := time.Now()
	startDate := time.Date(curTime.Year(), curTime.Month(), curTime.Day(), 00, 00, 00, 00, time.Local)
	endDate := time.Date(curTime.Year(), curTime.Month(), curTime.Day(), 23, 59, 59, 00, time.Local)

	for _, stock := range stocks {
		symbol := stock.StockCode
		switch pb.StockExchange(pb.StockExchange_value[stock.StockExchange]) {
		case pb.StockExchange_SSE:
			symbol += ".SS"
		case pb.StockExchange_SZSE:
			symbol += ".SZ"
		}

		log.Info("[SyncStockMinAtTodayJob]同步股票(%s)", symbol)
		params := &chart.Params{
			Symbol:   symbol,
			Interval: datetime.OneMin,
			Start:    datetime.New(&startDate),
			End:      datetime.New(&endDate),
		}
		iter := chart.Get(params)
		var stockPrices []*model.StockPriceItem
		for iter.Next() {
			b := iter.Bar()
			price := new(model.StockPriceItem)
			price.Time = time.Unix(int64(b.Timestamp), 0)
			price.High, _ = b.High.Float64()
			price.Volume = b.Volume
			price.Open, _ = b.Open.Float64()
			price.Close, _ = b.Close.Float64()
			price.Low, _ = b.Low.Float64()
			stockPrices = append(stockPrices, price)
		}
		if err = iter.Err(); err != nil {
			log.Error("[SyncStockMinAtTodayJob] symbol(%s) iter.Err(%+v)", symbol, err)
			continue
		}

		stockMin := &model.StockMin{
			Date:          curTime,
			StockExchange: stock.StockExchange,
			StockCode:     stock.StockCode,
			StockName:     stock.StockName,
			Prices:        stockPrices,
		}
		_, err := s.dao.AddStockMin(ctx, stockMin)
		if err != nil {
			log.Error("[SyncStockMinAtTodayJob] AddStockMin stockMin(%+v) err(%+v)", stockMin, err)
			return
		}
	}
	log.Info("SyncStockMinAtTodayJob end")
}
