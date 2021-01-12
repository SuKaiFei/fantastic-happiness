package service

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	pb "github.com/SuKaiFei/fantastic-happiness/api"
	"github.com/SuKaiFei/fantastic-happiness/internal/dao"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	kLog "github.com/go-kratos/kratos/pkg/log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
	"github.com/piquette/finance-go"
	"github.com/robfig/cron/v3"
)

var (
	Provider = wire.NewSet(New, wire.Bind(new(pb.StockCollectServer), new(*Service)))

	isTesting = false
)

type AppApi struct {
	Juhe struct {
		StockApiKey string
	}
}

// Service service.
type Service struct {
	ctx           context.Context
	ctxCancelFunc context.CancelFunc

	ac     *paladin.Map
	appApi *AppApi

	dao dao.Dao

	c *cron.Cron
}

//go:generate kratos tool wire
// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	s = &Service{
		ctx:           ctx,
		ctxCancelFunc: cancelFunc,
		ac:            &paladin.TOML{},
		dao:           d,
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

	finance.SetHTTPClient(
		&http.Client{
			Timeout: 1 * time.Minute,
		},
	)

	if !isTesting {
		if _, err = s.c.AddFunc("0 30 9 * * *", s.jobSyncSSEStockList); err != nil {
			return nil, nil, err
		}
		if _, err = s.c.AddFunc("0 30 9 * * *", s.jobSyncSZSEStockList); err != nil {
			return nil, nil, err
		}
		if _, err = s.c.AddFunc("0 0 16 * * *", func() {
			if _, err := s.SyncStockMinAtToday(s.ctx, new(empty.Empty)); err != nil {
				kLog.Error("SyncStockMinAtToday err(%+v)", err)
				return
			}
		}); err != nil {
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

	if s.ctxCancelFunc != nil {
		s.ctxCancelFunc()
	}
}
