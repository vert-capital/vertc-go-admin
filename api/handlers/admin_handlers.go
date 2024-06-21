package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	api "github.com/vert-capital/vertc-go-admin/api/entity"
	"github.com/vert-capital/vertc-go-admin/entity"
	"github.com/vert-capital/vertc-go-admin/infrasctructure/database/repository"
	usecases_tables "github.com/vert-capital/vertc-go-admin/usecases/usecase_tables"
	"gorm.io/gorm"
)

type AdminHandlers struct {
	UcTable usecases_tables.IUseCaseTable
	Tables  map[string]entity.Table
}

func NewAdminHandler(ucTable usecases_tables.IUseCaseTable) *AdminHandlers {
	return &AdminHandlers{
		UcTable: ucTable,
	}
}

func (ah *AdminHandlers) ListTables(c *gin.Context) {
	data := api.ResponseListTables{}
	for _, table := range ah.Tables {
		data[table.Category] = table
	}
	jsonResponse(c, 200, data)

}

func (ah *AdminHandlers) GetInfoTable(c *gin.Context) {
	tableName := c.Param("table_name")
	table := ah.Tables[tableName]
	if table.Name == "" {
		handleError(c, nil)
		return
	}
	jsonResponse(c, 200, table)
}

func (ah *AdminHandlers) ListTable(c *gin.Context) {
	tableName := c.Param("table_name")
	table := ah.Tables[tableName]
	if table.Name == "" {
		handleError(c, nil)
		return
	}
	filters := api.Filters{}
	search, exists := c.GetQuery("search")
	if exists {
		filters.Search = &search
	}
	page, exists := c.GetQuery("page")
	if exists {
		page, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			handleError(c, err)
			return
		}
		filters.Page = int(page)
	}
	pageSize, exists := c.GetQuery("page_size")
	if exists {
		pageSize, err := strconv.ParseInt(pageSize, 10, 64)
		if err != nil {
			handleError(c, err)
			return
		}
		filters.PageSize = int(pageSize)
	}
	fields := []api.FilterField{}
	for key, value := range c.Request.URL.Query() {
		if key != "search" && key != "page" && key != "page_size" {
			fields = append(fields, api.FilterField{
				Field: key,
				Value: value[0],
			})
		}
	}
	if len(fields) > 0 {
		filters.FilterFields = &fields
	}
	actions, exists := c.GetQuery("actions")
	if exists {
		response, err := runAction(table, actions)
		if err != nil {
			handleError(c, err)
			return
		}
		jsonResponse(c, 200, response)
		return
	}
	response, err := ah.UcTable.List(table, filters)
	if err != nil {
		handleError(c, err)
		return
	}
	jsonResponse(c, 200, response)
}

func (ah *AdminHandlers) GetTable(c *gin.Context) {
	tableName := c.Param("table_name")
	table := ah.Tables[tableName]
	if table.Name == "" {
		handleError(c, nil)
		return
	}
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		handleError(c, err)
		return

	}

	response, err := ah.UcTable.Get(table, id)
	if err != nil {
		handleError(c, err)
		return
	}
	jsonResponse(c, 200, response)
}

func (ah *AdminHandlers) CreateTable(c *gin.Context) {
	tableName := c.Param("table_name")
	table := ah.Tables[tableName]
	if table.Name == "" {
		handleError(c, nil)
		return
	}
	var data map[string]interface{}
	err := c.ShouldBindJSON(data)
	if err != nil {
		handleError(c, err)
		return

	}
	response, err := ah.UcTable.Create(table, data)
	if err != nil {
		handleError(c, err)
		return
	}
	jsonResponse(c, 200, response)
}

func (ah *AdminHandlers) UpdateTable(c *gin.Context) {
	tableName := c.Param("table_name")
	table := ah.Tables[tableName]
	if table.Name == "" {
		handleError(c, nil)
		return
	}
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		handleError(c, err)
		return
	}

	var data map[string]interface{}
	err = c.ShouldBindJSON(data)
	if err != nil {
		handleError(c, err)
		return
	}

	response, err := ah.UcTable.Update(table, data, id)
	if err != nil {
		handleError(c, err)
		return
	}
	jsonResponse(c, 200, response)
}

func (ah *AdminHandlers) DeleteTable(c *gin.Context) {
	type DeleteRequest struct {
		Ids []int `json:"ids"`
	}
	tableName := c.Param("table_name")
	table := ah.Tables[tableName]
	if table.Name == "" {
		handleError(c, nil)
		return
	}
	ids := &DeleteRequest{}
	err := c.ShouldBindJSON(ids)
	if err != nil {
		handleError(c, err)
		return
	}
	response, err := ah.UcTable.Delete(table, ids.Ids)
	if err != nil {
		handleError(c, err)
		return
	}
	jsonResponse(c, 200, response)
}

func MountTableHandlers(gin *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewRepositoryTable(db)
	uc := usecases_tables.NewUseCaseTable(repo)
	ah := NewAdminHandler(uc)
	gin.GET("/tables", ah.ListTables)
	gin.GET("/tables/:table_name", ah.GetInfoTable)
	gin.GET("/tables/:table_name/crud", ah.ListTable)

	gin.POST("/tables/:table_name/crud", ah.CreateTable)
	gin.DELETE("/tables/:table_name/crud", ah.DeleteTable)

	gin.GET("/tables/:table_name/crud/:id", ah.GetTable)
	gin.PUT("/tables/:table_name/crud/:id", ah.UpdateTable)
}
