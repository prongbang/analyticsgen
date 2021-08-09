package generate

import (
	"github.com/prongbang/analyticsgen/internal/analyticsgen/sheets"
	"github.com/prongbang/analyticsgen/pkg/csvx"
)

type AndroidUseCase interface {
	UseCase
}

type androidUc struct {
	SheetUc sheets.UseCase
}

func (a *androidUc) GenKey(sheet sheets.Sheets) (string, error) {
	panic("implement me")
}

func (a *androidUc) GenCode(sheet sheets.Sheets) (string, error) {
	panic("implement me")
}

func (a *androidUc) PrepareFunction(values csvx.CsvList) map[string][]interface{} {
	panic("implement me")
}

func (a *androidUc) BuildFunction(statement map[string]interface{}) string {
	panic("implement me")
}

func NewAndroidUseCase(sheetUc sheets.UseCase) AndroidUseCase {
	return &androidUc{
		SheetUc: sheetUc,
	}
}
