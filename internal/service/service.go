package service

import (
	"context"
	"log"
	"os"

	pb "github.com/SuKaiFei/fantastic-happiness/api"
	"github.com/SuKaiFei/fantastic-happiness/internal/dao"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
	"github.com/robfig/cron/v3"
)

var (
	Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)))

	isTesting = false
)

type AppApi struct {
	Juhe struct {
		StockApiKey string
	}
}

// Service service.
type Service struct {
	ac     *paladin.Map
	appApi *AppApi

	dao dao.Dao

	c *cron.Cron
}

//go:generate kratos tool wire
// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
		c: cron.New(
			cron.WithSeconds(),
			cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "", log.LstdFlags))),
		),
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	if err != nil {
		return nil, nil, err
	}

	s.appApi = new(AppApi)

	err = s.ac.Get("api").UnmarshalTOML(s.appApi)
	if err != nil {
		return nil, nil, err
	}

	if !isTesting {
		if _, err = s.c.AddFunc("0 30 9 * * *", s.jobSyncSSEStockList); err != nil {
			return nil, nil, err
		}
		if _, err = s.c.AddFunc("0 30 9 * * *", s.jobSyncSZSEStockList); err != nil {
			return nil, nil, err
		}
		if _, err = s.c.AddFunc("0 0 16 * * *", s.jobSyncSZSEStockList); err != nil {
			return nil, nil, err
		}
	}
	s.c.Start()

	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	<-s.c.Stop().Done()
}
