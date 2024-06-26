package vertc_go_admin

type ResponseListTables map[string]interface{}

type ResponseList struct {
	Page       int                      `json:"page"`
	PageSize   int                      `json:"page_size"`
	Total      int                      `json:"total"`
	TotalPages int                      `json:"total_pages"`
	Data       []map[string]interface{} `json:"data"`
}

type ResponseCreateUpdateDelete struct {
	Message string `json:"message"`
}

type FilterField struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type Filters struct {
	Search       *string        `json:"search"`
	FilterFields *[]FilterField `json:"fields"`
	OrderBy      []string       `json:"order_by"`
	Page         int            `json:"page"`
	PageSize     int            `json:"page_size"`
}
