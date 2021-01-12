package service

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestServiceSyncStockMinAtToday(t *testing.T) {
	Convey("SyncStockMinAtToday", t, func() {
		var (
			a2 = &emptypb.Empty{}
		)
		Convey("When everything goes positive", func() {
			p1, err := s.SyncStockMinAtToday(ctx, a2)
			Convey("Then err should be nil.p1 should not be nil.", func() {
				So(err, ShouldBeNil)
				So(p1, ShouldNotBeNil)
			})
		})
	})
}

func TestServiceSyncStockMinAtTodayJob(t *testing.T) {
	Convey("SyncStockMinAtTodayJob", t, func() {
		Convey("When everything goes positive", func() {
			s.SyncStockMinAtTodayJob()
			Convey("No return values", func() {
			})
		})
	})
}
