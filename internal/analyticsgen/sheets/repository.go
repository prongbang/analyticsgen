package sheets

import (
	"errors"
	"fmt"
	"github.com/prongbang/analyticsgen/pkg/csvx"
	"github.com/prongbang/callx"
)

type Repository interface {
	Get(sheet Sheets) (csvx.CsvList, error)
}

type repository struct {
	CallX callx.CallX
	CsvX  csvx.CsvX
}

func (r *repository) Get(sheet Sheets) (csvx.CsvList, error) {
	resp := r.CallX.Get(fmt.Sprintf("/%s/export?format=csv&id=%s&gid=%s", sheet.DocumentId, sheet.DocumentId, sheet.Id))
	if resp.Code == 200 {
		return r.CsvX.ReadAll(string(resp.Data)), nil
	}
	return csvx.CsvList{}, errors.New(fmt.Sprintf("Error code %d", resp.Code))
}

func NewRepository(callX callx.CallX, csvX csvx.CsvX) Repository {
	return &repository{
		CallX: callX,
		CsvX:  csvX,
	}
}
