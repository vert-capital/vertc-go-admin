package vertc_go_admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminHandlers struct {
	UcTable IUseCaseTable
}

func NewAdminHandler(ucTable IUseCaseTable) *AdminHandlers {
	return &AdminHandlers{
		UcTable: ucTable,
	}
}

func (ah *AdminHandlers) ListTables(c *gin.Context) {
	data := ResponseListTables{}
	for _, table := range Tabelas {
		table.Fields = nil
		data[table.Category] = table
	}
	jsonResponse(c, 200, data)

}

func (ah *AdminHandlers) GetInfoTable(c *gin.Context) {
	tableName := c.Param("table_name")
	table := Tabelas[tableName]
	if table.Name == "" {
		handleError(c, nil)
		return
	}
	jsonResponse(c, 200, table)
}

func (ah *AdminHandlers) ListTable(c *gin.Context) {
	tableName := c.Param("table_name")
	table := Tabelas[tableName]
	if table.Name == "" {
		handleError(c, nil)
		return
	}
	filters := Filters{}
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
	fields := []FilterField{}
	for key, value := range c.Request.URL.Query() {
		if key != "search" && key != "page" && key != "page_size" {
			fields = append(fields, FilterField{
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
	table := Tabelas[tableName]
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
	table := Tabelas[tableName]
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
	table := Tabelas[tableName]
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
	table := Tabelas[tableName]
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
	repo := NewRepositoryTable(db)
	uc := NewUseCaseTable(repo)
	ah := NewAdminHandler(uc)
	gin.GET("/tables", ah.ListTables)
	gin.GET("/tables/:table_name", ah.GetInfoTable)
	gin.GET("/tables/:table_name/crud", ah.ListTable)

	gin.POST("/tables/:table_name/crud", ah.CreateTable)
	gin.DELETE("/tables/:table_name/crud", ah.DeleteTable)

	gin.GET("/tables/:table_name/crud/:id", ah.GetTable)
	gin.PUT("/tables/:table_name/crud/:id", ah.UpdateTable)
}
