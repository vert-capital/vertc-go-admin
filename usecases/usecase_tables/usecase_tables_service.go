package usecases_tables

import (
	api "github.com/vert-capital/vertc-go-admin/api/entity"
	"github.com/vert-capital/vertc-go-admin/entity"
)

type UseCaseTable struct {
	repo IRepositoryTable
}

func NewUseCaseTable(repo IRepositoryTable) *UseCaseTable {
	return &UseCaseTable{
		repo: repo,
	}
}

func (uc *UseCaseTable) List(table entity.Table, filters api.Filters) (response api.ResponseList, err error) {
	return uc.repo.List(table, filters)
}

func (uc *UseCaseTable) Get(table entity.Table, id int) (response entity.Fields, err error) {
	return uc.repo.Get(table, id)
}

func (uc *UseCaseTable) Create(table entity.Table, fields entity.Fields) (response api.ResponseCreateUpdateDelete, err error) {
	return uc.repo.Create(table, fields)
}

func (uc *UseCaseTable) Update(table entity.Table, fields entity.Fields, id int) (response api.ResponseCreateUpdateDelete, err error) {
	return uc.repo.Update(table, fields, id)
}

func (uc *UseCaseTable) Delete(table entity.Table, ids []int) (response api.ResponseCreateUpdateDelete, err error) {
	return uc.repo.Delete(table, ids)
}
