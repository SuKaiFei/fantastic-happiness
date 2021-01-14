package service

import (
	"context"
	"fmt"
	pb "github.com/SuKaiFei/fantastic-happiness/api"
	"github.com/SuKaiFei/fantastic-happiness/internal/model"
	"github.com/SuKaiFei/fantastic-happiness/pkg/juhe"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

// GetStock 获取股票信息
func (s *Service) GetStock(ctx context.Context, exchange pb.StockExchange, stockCode string) (*empty.Empty, error) {
	stockResp, err := juhe.GetHSFinanceStock(s.appApi.Juhe.StockApiKey, convertJuHeGid(exchange, stockCode))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	println(stockResp)

	return &empty.Empty{}, nil
}

// GetSZStockList 获取深圳交易所股票列表信息
func (s *Service) GetSZSEStockList(ctx context.Context) (stocks []*model.Stock, err error) {
	var (
		page    uint32
		curTime = time.Now()
	)

	for {
		select {
		case <-ctx.Done():
			return nil, errors.WithStack(ctx.Err())
		default:
		}

		page++

		stockResp, err := juhe.GetSzAllFinanceStock(s.appApi.Juhe.StockApiKey, page)
		if err != nil {
			return nil, errors.Wrapf(err, "page(%d)", page)
		}
		if len(stockResp.Result.Data) == 0 {
			break
		}

		stocksTmp := make([]*model.Stock, len(stockResp.Result.Data))
		for i, datum := range stockResp.Result.Data {
			priceTime, _ := time.ParseInLocation(
				"2006-01-02 15:04:05",
				fmt.Sprintf("%04d-%02d-%02d %s", curTime.Year(), curTime.Month(), curTime.Day(), datum.Ticktime),
				time.Local,
			)
			price, _ := strconv.ParseFloat(datum.Trade, 64)
			stocksTmp[i] = &model.Stock{
				Date:          curTime,
				StockExchange: pb.StockExchange_SZSE.String(),
				StockCode:     datum.Code,
				StockName:     datum.Name,
				Prices: []*model.StockPrice{
					{
						Time:  priceTime,
						Price: price,
					},
				},
				MktCap:        datum.Mktcap,
				Nmc:           datum.Nmc,
				TurnoverRatio: datum.Turnoverratio,
				PBRatio:       datum.Pb,
			}
		}
		stocks = append(stocks, stocksTmp...)
	}

	return
}

// GetSSEStockList 获取上海交易所股票列表信息
func (s *Service) GetSSEStockList(ctx context.Context) (stocks []*model.Stock, err error) {
	var (
		page    uint32
		curTime = time.Now()
	)

	for {
		select {
		case <-ctx.Done():
			return nil, errors.WithStack(ctx.Err())
		default:
		}

		page++

		stockResp, err := juhe.GetShAllFinanceStock(s.appApi.Juhe.StockApiKey, page)
		if err != nil {
			return nil, errors.Wrapf(err, "page(%d)", page)
		}
		if len(stockResp.Result.Data) == 0 {
			break
		}

		stocksTmp := make([]*model.Stock, len(stockResp.Result.Data))
		for i, datum := range stockResp.Result.Data {
			priceTime, _ := time.ParseInLocation(
				"2006-01-02 15:04:05",
				fmt.Sprintf("%04d-%02d-%02d %s", curTime.Year(), curTime.Month(), curTime.Day(), datum.Ticktime),
				time.Local,
			)
			price, _ := strconv.ParseFloat(datum.Trade, 64)
			stocksTmp[i] = &model.Stock{
				Date:          curTime,
				StockExchange: pb.StockExchange_SSE.String(),
				StockCode:     datum.Code,
				StockName:     datum.Name,
				Prices: []*model.StockPrice{
					{
						Time:  priceTime,
						Price: price,
					},
				},
				MktCap:        datum.Mktcap,
				Nmc:           datum.Nmc,
				TurnoverRatio: datum.Turnoverratio,
				PBRatio:       datum.Pb,
			}
		}
		stocks = append(stocks, stocksTmp...)
	}

	return
}

func convertJuHeGid(exchange pb.StockExchange, stockCode string) string {
	switch exchange {
	case pb.StockExchange_SSE:
		return "sh" + stockCode
	case pb.StockExchange_SZSE:
		return "sz" + stockCode
	}

	return stockCode
}
