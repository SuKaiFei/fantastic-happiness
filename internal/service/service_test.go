package service

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"os"
	"testing"
)

var (
	ctx = context.Background()
	s   *Service

	_ = func() interface{} {
		isTesting = true
		return nil
	}
)

func TestMain(m *testing.M) {
	isTesting = true
	_ = flag.Set("conf", "../../configs")
	flag.Parse()
	_ = paladin.Init()

	var (
		cf  func()
		err error
	)

	if s, cf, err = newTestService(); err != nil {
		panic(err)
	}

	ret := m.Run()
	cf()

	os.Exit(ret)
}
