package service

import (
	"testing"

	pb "github.com/SuKaiFei/fantastic-happiness/api"
	. "github.com/smartystreets/goconvey/convey"
)

func TestServiceGetStock(t *testing.T) {
	Convey("GetStock", t, func() {
		var (
			exchange  = pb.StockExchange_SZSE
			stockCode = "300001"
			//exchange  = pb.StockExchange_SSE
			//stockCode = "600004"
		)
		Convey("When everything goes positive", func() {
			p1, err := s.GetStock(ctx, exchange, stockCode)
			Convey("Then err should be nil.p1 should not be nil.", func() {
				So(err, ShouldBeNil)
				So(p1, ShouldNotBeNil)
			})
		})
	})
}

func TestServiceGetSSEStockList(t *testing.T) {
	Convey("GetSSEStockList", t, func() {
		var ()
		Convey("When everything goes positive", func() {
			p1, err := s.GetSSEStockList(ctx)
			if err == nil {
				for _, v := range p1 {
					t.Logf("%+v", v)
				}
			}
			Convey("Then err should be nil.p1 should not be nil.", func() {
				So(err, ShouldBeNil)
				So(p1, ShouldNotBeNil)
			})
		})
	})
}

func TestServiceGetSZSEStockList(t *testing.T) {
	Convey("GetSZSEStockList", t, func() {
		var ()
		Convey("When everything goes positive", func() {
			p1, err := s.GetSZSEStockList(ctx)
			if err == nil {
				for _, v := range p1 {
					t.Logf("%+v", v)
				}
			}
			Convey("Then err should be nil.p1 should not be nil.", func() {
				So(err, ShouldBeNil)
				So(p1, ShouldNotBeNil)
			})
		})
	})
}
