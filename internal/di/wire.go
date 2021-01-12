// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/SuKaiFei/fantastic-happiness/internal/dao"
	"github.com/SuKaiFei/fantastic-happiness/internal/server/grpc"
	"github.com/SuKaiFei/fantastic-happiness/internal/server/http"
	"github.com/SuKaiFei/fantastic-happiness/internal/service"
	colSvc "github.com/SuKaiFei/fantastic-happiness/internal/service/collect"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, colSvc.Provider, http.New, grpc.New, NewApp))
}
