//+build wireinject

package sheets

import (
	"github.com/google/wire"
	"github.com/prongbang/analyticsgen/pkg/csvx"
)

func New() UseCase {
	wire.Build(
		NewRepository,
		NewUseCase,
		NewCallX,
		csvx.New,
	)
	return nil
}
