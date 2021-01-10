package service

import (
	"context"
	"github.com/go-kratos/kratos/pkg/log"
)

// jobSyncSSEStockList 作业同步今天上海证券交易所股票列表
func (s *Service) jobSyncSSEStockList() {
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

	return
}

// jobSyncSZSEStockList 作业同步今天深圳证券交易所股票列表
func (s *Service) jobSyncSZSEStockList() {
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

	return
}
