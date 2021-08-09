package generate

import (
	"github.com/prongbang/analyticsgen/internal/analyticsgen/sheets"
	"github.com/prongbang/analyticsgen/pkg/csvx"
)

type UseCase interface {
	GenKey(sheet sheets.Sheets) (string, error)
	GenCode(sheet sheets.Sheets) (string, error)
	PrepareFunction(values csvx.CsvList) map[string][]interface{}
	BuildFunction(statement map[string]interface{}) string
}
