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
				StockExchange: pb.StockExchange_SSE,
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
