package service

import (
	"context"
	"github.com/go-kratos/kratos/pkg/log"
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

	log.Info("jobSyncSZSEStockList start")
	return
}
