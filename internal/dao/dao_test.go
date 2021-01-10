package dao

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
)

var d *dao
var ctx = context.Background()

func TestMain(m *testing.M) {
	_ = flag.Set("conf", "../../configs")
	flag.Parse()
	var err error
	if err = paladin.Init(); err != nil {
		panic(err)
	}
	var cf func()
	if d, cf, err = newTestDao(); err != nil {
		panic(err)
	}
	ret := m.Run()
	cf()
	os.Exit(ret)
}
