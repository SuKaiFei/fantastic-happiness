package service

import (
	"context"
	pb "github.com/SuKaiFei/fantastic-happiness/api"
	"github.com/SuKaiFei/fantastic-happiness/internal/model"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"time"
)

// jobSyncSSEStockList 作业同步今天上海证券交易所股票列表
func (s *Service) jobSyncSSEStockList() {
	log.Info("jobSyncSSEStockList start")
	ctx := context.Background()

	stocks, err := s.GetSSEStockList(ctx)
	if err != nil {
		log.Error("GetSSEStockList err(%+v)", err)
		return
	}

	for _, stock := range stocks {
		_, err := s.dao.AddStock(ctx, stock)
		if err != nil {
			log.Error("AddStock stock(%+v) err(%+v)", stock, err)
			return
		}
	}

	log.Info("jobSyncSSEStockList end")
	return
}

// jobSyncSZSEStockList 作业同步今天深圳证券交易所股票列表
func (s *Service) jobSyncSZSEStockList() {
	log.Info("jobSyncSZSEStockList start")
	ctx := context.Background()

	stocks, err := s.GetSZSEStockList(ctx)
	if err != nil {
		log.Error("GetSZSEStockList err(%+v)", err)
		return
	}

	for _, stock := range stocks {
		_, err := s.dao.AddStock(ctx, stock)
		if err != nil {
			log.Error("AddStock stock(%+v) err(%+v)", stock, err)
			return
		}
	}

	log.Info("jobSyncSZSEStockList end")
	return
}

// jobSyncStockToday 作业同步今天深沪证券交易所股票
func (s *Service) jobSyncStockToday() {
	log.Info("jobSyncStockToday start")
	ctx := context.Background()

	stocks, err := s.dao.GetStockByToday(ctx, "_id", "stock_exchange", "stock_name", "stock_code")
	if err != nil {
		log.Error("GetStockByToday err(%+v)", err)
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
			log.Error("symbol(%s) iter.Err(%+v)", symbol, err)
			continue
		}

		stockDay := &model.StockDay{
			Date:          curTime,
			StockExchange: stock.StockExchange,
			StockCode:     stock.StockCode,
			StockName:     stock.StockName,
			Prices:        stockPrices,
		}
		_, err := s.dao.AddStockDay(ctx, stockDay)
		if err != nil {
			log.Error("AddStockDay stockDay(%+v) err(%+v)", stockDay, err)
			return
		}

	}
	log.Info("jobSyncStockToday end")
	return
}
