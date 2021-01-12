// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package service

import (
	"github.com/SuKaiFei/fantastic-happiness/internal/dao"
	"github.com/google/wire"
)

func newTestService() (*Service, func(), error) {
	panic(wire.Build(Provider, dao.Provider))
}
