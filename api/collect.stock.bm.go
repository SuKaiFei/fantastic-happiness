// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: collect.stock.proto

package api

import (
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"
)
import google_protobuf1 "google.golang.org/protobuf/types/known/emptypb"

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathStockCollectPing = "/github.com.SuKaiFei.fantastic_happiness.service.v1.StockCollect/Ping"
var PathStockCollectSyncStockMinAtToday = "/api/v1/stock_collect/sync/stock_min/today"

// StockCollectBMServer is the server API for StockCollect service.
type StockCollectBMServer interface {
	Ping(ctx context.Context, req *google_protobuf1.Empty) (resp *google_protobuf1.Empty, err error)

	// 作业同步今天深沪证券交易所股票-分钟数据
	SyncStockMinAtToday(ctx context.Context, req *google_protobuf1.Empty) (resp *google_protobuf1.Empty, err error)
}

var StockCollectSvc StockCollectBMServer

func stockCollectPing(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := StockCollectSvc.Ping(c, p)
	c.JSON(resp, err)
}

func stockCollectSyncStockMinAtToday(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := StockCollectSvc.SyncStockMinAtToday(c, p)
	c.JSON(resp, err)
}

// RegisterStockCollectBMServer Register the blademaster route
func RegisterStockCollectBMServer(e *bm.Engine, server StockCollectBMServer) {
	StockCollectSvc = server
	e.GET("/github.com.SuKaiFei.fantastic_happiness.service.v1.StockCollect/Ping", stockCollectPing)
	e.POST("/api/v1/stock_collect/sync/stock_min/today", stockCollectSyncStockMinAtToday)
}
