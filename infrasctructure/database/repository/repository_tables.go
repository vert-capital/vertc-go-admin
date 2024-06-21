package vertc_go_admin

import (
	api "github.com/vert-capital/vertc-go-admin/api/entity"
	entity "github.com/vert-capital/vertc-go-admin/entity"
	"gorm.io/gorm"
)

type RepositoryTable struct {
	DB *gorm.DB
}

func NewRepositoryTable(db *gorm.DB) *RepositoryTable {
	return &RepositoryTable{DB: db}
}

func (r RepositoryTable) List(table entity.Table, filters api.Filters) (response api.ResponseList, err error) {
	query := r.DB.Table(table.Name)
	if filters.Search != nil {
		for _, field := range table.SearchFields {
			query.Or(field+" LIKE ?", "%"+*filters.Search+"%")
		}
	}
	if filters.FilterFields != nil {
		for _, field := range *filters.FilterFields {
			query.Where(field.Field+" = ?", field.Value)
		}
	}
	if filters.OrderBy != nil {
		orderby := ""
		for idx, field := range filters.OrderBy {
			if field[0] == '-' {
				orderby += field[1:] + " DESC"
			} else {
				orderby += field + ""
			}
			if idx != len(filters.OrderBy)-1 {
				orderby += ", "
			}
		}
		query.Order(orderby)
	}
	cpquery := query
	err = query.Find(&response.Data).Error
	if err != nil {
		return api.ResponseList{
			Page:       0,
			PageSize:   0,
			TotalPages: 0,
			Total:      0,
			Data:       nil,
		}, err
	}
	response.Page = filters.Page
	response.PageSize = filters.PageSize
	response.Total = len(response.Data)
	response.TotalPages = response.Total / filters.PageSize
	if response.Total%filters.PageSize != 0 {
		response.TotalPages++
	}
	cpquery.Offset((filters.Page - 1) * filters.PageSize).Limit(filters.PageSize)
	err = cpquery.Find(&response.Data).Error
	if err != nil {
		return api.ResponseList{
			Page:       0,
			PageSize:   0,
			TotalPages: 0,
			Total:      0,
			Data:       nil,
		}, err
	}

	return response, nil
}

func (r RepositoryTable) Get(table entity.Table, id int) (response entity.Fields, err error) {
	err = r.DB.Table(table.Name).Where("id = ?", id).First(&response).Error
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r RepositoryTable) Create(table entity.Table, fields entity.Fields) (response api.ResponseCreateUpdateDelete, err error) {
	err = r.DB.Table(table.Name).Create(fields).Error
	if err != nil {
		return api.ResponseCreateUpdateDelete{
			Message: "Error creating record",
		}, err
	}
	return api.ResponseCreateUpdateDelete{
		Message: "Record created successfully",
	}, nil
}

func (r RepositoryTable) Update(table entity.Table, fields entity.Fields, id int) (response api.ResponseCreateUpdateDelete, err error) {
	err = r.DB.Table(table.Name).Where("id = ?", id).Updates(fields).Error
	if err != nil {
		return api.ResponseCreateUpdateDelete{
			Message: "Error updating record",
		}, err
	}
	return api.ResponseCreateUpdateDelete{
		Message: "Record updated successfully",
	}, nil
}

func (r RepositoryTable) Delete(table entity.Table, ids []int) (response api.ResponseCreateUpdateDelete, err error) {
	err = r.DB.Table(table.Name).Where("id IN (?)", ids).Delete(nil).Error
	if err != nil {
		return api.ResponseCreateUpdateDelete{
			Message: "Error deleting record",
		}, err
	}
	return api.ResponseCreateUpdateDelete{
		Message: "Record deleted successfully",
	}, nil
}
