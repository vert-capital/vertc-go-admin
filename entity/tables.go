package entity

type Fields map[string]interface{}

type Actions map[string](func() error)

type Table struct {
	Name         string    `json:"name"`
	Category     string    `json:"category"`
	Fields       Fields    `json:"fields"`
	SearchFields []*string `json:"search_fields"`
	Actions      Actions   `json:"actions"`
}
