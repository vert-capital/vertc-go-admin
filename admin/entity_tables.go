package vertc_go_admin

type Fields map[string]interface{}

type Actions map[string](func() error)

type Tables map[string]Table

type Table struct {
	Name         string   `json:"name"`
	Category     string   `json:"category"`
	Fields       *Fields  `json:"fields"`
	SearchFields []string `json:"search_fields"`
	Actions      *Actions `json:"actions"`
}

func AddTabela(tabela Table) {
	Tabelas[tabela.Name] = tabela
}
