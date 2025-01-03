# Documentação

## Instalação
<code>
go get github.com/vert-capital/vertc-go-admin <br>
go get github.com/vert-capital/vertc-go-admin/admin
</code>

## Objetivo
Ser uma biblioteca que aplicações da VERT usando GORM poderão criar automaticamente rotas para o ADMIN

## Implementação
main.go
```
import vertc_go_admin "github.com/vert-capital/vertc-go-admin/admin"

func main() {
	go kafka.StartKafka()
	go vertc_go_admin.RunServer(conn)
}
```


api.go
```

import vertc_go_admin "github.com/vert-capital/vertc-go-admin/admin"

func setupRouter(conn *gorm.DB) *gin.Engine {
	r := gin.New()
	vertc_go_admin.RoutersAdmin(r, conn, middleware)
}
```

Crie arquivos NO SEU PROJETO numa pasta admin/ para centralizar, seguindo o exemplo:

admin.go
```
func CreateAdminTables() {
	admin.EmissionAdmin()
}
```

admin_emission.go
```

func ActionPrintAllID(ids []string) error {
	db := postgres.Connect()
	repo := repository.NewEmissionPostgres(db)
	emissions, err := repo.GetEmissionList() 
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, emission := range emissions {
		fmt.Println(emission.ID)
	}
	return nil
}

func EmissionAdmin() {
	search_fields := []string{
		"number",
	}
	actions := vertc_go_admin.Actions{}
	actions["print_all_id"] = ActionPrintAllEmissionID
	// vertc_go_admin.SetAdmin(ENTIDADE, "NOME DA TABELA", "CATEGORIA PARA O MENU", "CAMPOS DE BUSCA", "AÇÕES")
	vertc_go_admin.SetAdmin(entity.EntityEmission{}, "entity_emissions", "core", search_fields, actions)

}
```


## Rotas
```

    URL: api/admin/v1/tables
    Listagem de tabelas
    Método: GET
    Exemplo response:
    {
        "Core": [
            {
                "name": "entity_emissions",
                "category": "Core",
                "fields": null,
                "search_fields": [
                    "number",
                    "emission_name",
                    "emission_code_name"
                ],
                "actions": null
            },
            {
                "name": "entity_patrimonies",
                "category": "Core",
                "fields": null,
                "search_fields": [],
                "actions": null
            }
        ],
        "Series": [
            {
                "name": "entity_series",
                "category": "Series",
                "fields": null,
                "search_fields": [
                    "code_cetip",
                    "code_isin"
                ],
                "actions": [
                    "print_all_series_id"
                ]
            },
            {
                "name": "entity_remunerations",
                "category": "Series",
                "fields": null,
                "search_fields": [
                    "name"
                ],
                "actions": null
            }
        ]
    }
```

<br>

```

    URL: api/admin/v1/tables/<:table_name:>
    Estrutura de dados da tabela selecionada 
    Método: GET 
    Exemplo response: 
    {
        "name": "entity_remunerations",
        "category": "Series",
        "fields": {
            "Convention": {
                "json_field": "convention",
                "required": true,
                "type": "string"
            },
            "CreatedAt": {
                "json_field": "created_at",
                "required": false,
                "type": "struct"
            },
            "ExternalID": {
                "json_field": "-",
                "required": false,
                "type": "int"
            },
            "ID": {
                "json_field": "id",
                "required": false,
                "type": "int"
            },
            "InterestCorrection": {
                "json_field": "interest_correction",
                "required": true,
                "type": "string"
            },
            "Name": {
                "json_field": "name",
                "required": true,
                "type": "string"
            },
            "Series": {
                "json_field": "serie",
                "required": false,
                "type": "struct"
            },
            "SeriesID": {
                "json_field": "serie_id",
                "required": true,
                "type": "int"
            },
            "UpdatedAt": {
                "json_field": "updated_at",
                "required": false,
                "type": "struct"
            }
        },
        "search_fields": [
            "name"
        ],
        "actions": null
    }
```
<br>

```
    URL: api/admin/v1/tables/<:table_name:>/crud
    Listagem de dados da tabela selecionada

    ?NOME_DO_CAMPO=XXX -> Filtragem
    ?search=xx -> Procura por campos pré-selecionados
    ?order_by=xx OR -xx -> Ordenação
    ?page = 1 -> Paginação
    ?page_size=10 -> tamanho da pagina
    ?actions=xx -> Chamando ACTION
    
    Método: GET
    Exemplo response:
    {
        "page": 1,
        "page_size": 10,
        "total": 1,
        "total_pages": 1,
        "data": [
            {
                "convention": "teste",
                "created_at": null,
                "external_id": null,
                "id": 2,
                "interest_correction": "teste",
                "name": "%TESTE",
                "series_id": 4514,
                "updated_at": null
            }
        ]
    }

    Exemplo response (quando chamada a action): 
    {
	"message": "Action executed"
    }
```

<br>

```
    URL: api/admin/v1/tables/<:table_name:>/crud/<:id:>
    Get da tabela selecionada
    Método: GET/PUT 
    Exemplo response (GET):
    {
        "convention": "teste",
        "created_at": null,
        "external_id": null,
        "id": 2,
        "interest_correction": "teste",
        "name": "%TESTE",
        "series_id": 4514,
        "updated_at": null
    }
    Exemplo response (PUT):
    {
        "message": "Record created successfully"
    }

```

<br>

```
    URL: api/admin/v1/tables/<:table_name:>/crud/ 
    CRUD da tabela selecionada
    Método: DELETE/POST 
    Exemplo response (POST): 
    {
        "message": "Record created successfully"
    }
    Exemploy Payload (DELETE)
    {
        "ids": ["1","2","3","4"]
    }
    
    Exemplo response (DELETE): 
    {
        "message": "Record deleted successfully"
    }
```

<br>

```
    URL: api/admin/v1/tables/<:table_name:>/crud/?id= 
    CRUD da tabela selecionada 
    Método: DELETE 
    Exemplo response (DELETE): 
    {
	    "message": "Record deleted successfully"
    }

```
