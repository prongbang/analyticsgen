//+build wireinject

package analyticsgen

import (
	"github.com/google/wire"
	"github.com/prongbang/analyticsgen/internal/analyticsgen/generate"
	"github.com/prongbang/analyticsgen/internal/analyticsgen/sheets"
	"github.com/prongbang/filex"
)

func New() AnalyticsGen {
	wire.Build(
		filex.New,
		sheets.New,
		generate.ProviderSet,
		NewAnalyticsGen,
	)
	return nil
}
