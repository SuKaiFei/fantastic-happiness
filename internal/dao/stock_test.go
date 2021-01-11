package dao

import (
	pb "github.com/SuKaiFei/fantastic-happiness/api"
	"testing"
	"time"

	"github.com/SuKaiFei/fantastic-happiness/internal/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDaoAddStock(t *testing.T) {
	Convey("AddStock", t, func() {
		var (
			stock = &model.Stock{
				Date:          time.Now(),
				StockExchange: pb.StockExchange_SSE.String(),
				StockCode:     "su",
				StockName:     "test",
				Prices:        []*model.StockPrice{{Time: time.Now(), Price: 6666}},
				MktCap:        1.1,
				Nmc:           2.2,
				TurnoverRatio: 3.3,
				PBRatio:       4.4,
			}
		)
		Convey("When everything goes positive", func() {
			result, err := d.AddStock(ctx, stock)
			Convey("Then err should be nil.result should not be nil.", func() {
				So(err, ShouldBeNil)
				So(result, ShouldNotBeNil)
			})
		})
	})
}

func TestDaoGetStockByToday(t *testing.T) {
	Convey("GetStockByToday", t, func() {
		var (
			projections = []string{"_id"}
		)
		Convey("When everything goes positive", func() {
			stocks, err := d.GetStockByToday(ctx, projections...)
			if err == nil {
				for _, stock := range stocks {
					t.Logf("%+v", stock)
				}
			}
			Convey("Then err should be nil.result should not be nil.", func() {
				So(err, ShouldBeNil)
				So(stocks, ShouldNotBeNil)
			})
		})
	})
}
