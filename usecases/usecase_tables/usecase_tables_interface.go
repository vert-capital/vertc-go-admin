package vertc_go_admin

import (
	api "github.com/vert-capital/vertc-go-admin/api/entity"
	entity "github.com/vert-capital/vertc-go-admin/entity"
)

type IRepositoryTable interface {
	List(table entity.Table, filters api.Filters) (response api.ResponseList, err error)
	Get(table entity.Table, id int) (response entity.Fields, err error)
	Create(table entity.Table, fields entity.Fields) (response api.ResponseCreateUpdateDelete, err error)
	Update(table entity.Table, fields entity.Fields, id int) (response api.ResponseCreateUpdateDelete, err error)
	Delete(table entity.Table, ids []int) (response api.ResponseCreateUpdateDelete, err error)
}

type IUseCaseTable interface {
	List(table entity.Table, filters api.Filters) (response api.ResponseList, err error)
	Get(table entity.Table, id int) (response entity.Fields, err error)
	Create(table entity.Table, fields entity.Fields) (response api.ResponseCreateUpdateDelete, err error)
	Update(table entity.Table, fields entity.Fields, id int) (response api.ResponseCreateUpdateDelete, err error)
	Delete(table entity.Table, ids []int) (response api.ResponseCreateUpdateDelete, err error)
}
