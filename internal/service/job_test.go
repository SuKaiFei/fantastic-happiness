package service

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServicejobSyncSSStockList(t *testing.T) {
	Convey("jobSyncSSEStockList", t, func() {
		Convey("When everything goes positive", func() {
			s.jobSyncSSEStockList()
			Convey("No return values", func() {
			})
		})
	})
}
