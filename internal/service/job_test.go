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

func TestServicejobSyncSZSEStockList(t *testing.T) {
	Convey("jobSyncSZSEStockList", t, func() {
		Convey("When everything goes positive", func() {
			s.jobSyncSZSEStockList()
			Convey("No return values", func() {
			})
		})
	})
}

func TestServicejobSyncStockToday(t *testing.T) {
	Convey("jobSyncStockToday", t, func() {
		Convey("When everything goes positive", func() {
			s.jobSyncStockToday()
			Convey("No return values", func() {
			})
		})
	})
}
