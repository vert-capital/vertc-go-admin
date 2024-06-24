package vertc_go_admin

import (
	"gorm.io/gorm"
)

type RepositoryTable struct {
	DB *gorm.DB
}

func NewRepositoryTable(db *gorm.DB) *RepositoryTable {
	return &RepositoryTable{DB: db}
}

func (r RepositoryTable) List(table Table, filters Filters) (response ResponseList, err error) {
	var results []map[string]interface{}
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
	err = query.Find(&results).Error
	response.Data = results
	results = make([]map[string]interface{}, 0)
	if err != nil {
		return ResponseList{
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
	query.Offset((filters.Page - 1) * filters.PageSize).Limit(filters.PageSize)
	err = query.Find(&results).Error
	if err != nil {
		return ResponseList{
			Page:       0,
			PageSize:   0,
			TotalPages: 0,
			Total:      0,
			Data:       nil,
		}, err
	}
	response.Data = results
	return response, nil
}

func (r RepositoryTable) Get(table Table, id int) (response []map[string]interface{}, err error) {
	err = r.DB.Table(table.Name).Where("id = ?", id).Order("id desc").First(&response).Error
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r RepositoryTable) Create(table Table, fields Fields) (response ResponseCreateUpdateDelete, err error) {
	err = r.DB.Table(table.Name).Create(fields).Error
	if err != nil {
		return ResponseCreateUpdateDelete{
			Message: "Error creating record",
		}, err
	}
	return ResponseCreateUpdateDelete{
		Message: "Record created successfully",
	}, nil
}

func (r RepositoryTable) Update(table Table, fields Fields, id int) (response ResponseCreateUpdateDelete, err error) {
	err = r.DB.Table(table.Name).Where("id = ?", id).Updates(fields).Error
	if err != nil {
		return ResponseCreateUpdateDelete{
			Message: "Error updating record",
		}, err
	}
	return ResponseCreateUpdateDelete{
		Message: "Record updated successfully",
	}, nil
}

func (r RepositoryTable) Delete(table Table, ids []int) (response ResponseCreateUpdateDelete, err error) {
	err = r.DB.Table(table.Name).Where("id IN (?)", ids).Delete(nil).Error
	if err != nil {
		return ResponseCreateUpdateDelete{
			Message: "Error deleting record",
		}, err
	}
	return ResponseCreateUpdateDelete{
		Message: "Record deleted successfully",
	}, nil
}
