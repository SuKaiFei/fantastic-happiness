package service

import (
	"context"
	"github.com/SuKaiFei/fantastic-happiness/internal/service/common"
	"github.com/piquette/finance-go"
	"net/http"
	"time"

	pb "github.com/SuKaiFei/fantastic-happiness/api"
	"github.com/SuKaiFei/fantastic-happiness/internal/dao"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var (
	Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)))

	isTesting = false
)

// Service service.
type Service struct {
	ctx           context.Context
	ctxCancelFunc context.CancelFunc

	ac     *paladin.Map
	appApi *common.AppApi

	dao dao.Dao
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
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	if err != nil {
		return nil, nil, err
	}

	s.appApi = new(common.AppApi)

	err = s.ac.Get("api").UnmarshalTOML(s.appApi)
	if err != nil {
		return nil, nil, err
	}

	finance.SetHTTPClient(
		&http.Client{
			Timeout: 1 * time.Minute,
		},
	)

	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	if s.ctxCancelFunc != nil {
		s.ctxCancelFunc()
	}
}
