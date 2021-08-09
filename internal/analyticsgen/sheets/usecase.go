package sheets

import "github.com/prongbang/analyticsgen/pkg/csvx"

type UseCase interface {
	Get(sheet Sheets) csvx.CsvList
}

type useCase struct {
	Repos Repository
}

func (u *useCase) Get(sheet Sheets) csvx.CsvList {
	data, err := u.Repos.Get(sheet)
	if err != nil {
		return csvx.CsvList{}
	}
	return data
}

func NewUseCase(repos Repository) UseCase {
	return &useCase{
		Repos: repos,
	}
}
