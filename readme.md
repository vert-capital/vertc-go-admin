# Documentação

## Instalação
<code>
go get github.com/vert-capital/vertc-go-admin <br>
go get github.com/vert-capital/vertc-go-admin/admin
</code>

## Objetivo
Ser uma biblioteca que aplicações da VERT usando GORM poderão criar automaticamente rotas para o ADMIN

## Rotas
<code> 
    URL: api/admin/v1/tables <br>
    Listagem de tabelas <br>
    Método: GET <br>
    Exemplo response: <br>
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
    }<br>
</code>
<br>
<code> 
    URL: api/admin/v1/tables/<:table_name:> <br>
    Estrutura de dados da tabela selecionada <br>
    Método: GET <br>
    Exemplo response: <br>
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
    }<br>

</code>
<br>
<code> 
    URL: api/admin/v1/tables/<:table_name:>/crud <br>
    Listagem de dados da tabela selecionada<br>
    <br>
    ?NOME_DO_CAMPO=XXX -> Filtragem <br>
    ?search=xx -> Procura por campos pré-selecionados <br>
    ?order_by=xx OR -xx -> Ordenação <br>
    ?page = 1 -> Paginação <br>
    ?page_size=10 -> tamanho da pagina<br>
    ?actions=xx -> Chamando ACTION <br>
    <br>
    Método: GET <br>
    Exemplo response: <br>
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
    }<br>
    Exemplo response (quando chamada a action): <br>
    {
	"message": "Action executed"
}
</code>
<br>
<code> 
    URL: api/admin/v1/tables/<:table_name:>/crud/<:id:> <br>
    Get da tabela selecionada <br>
    Método: GET/PUT <br>
    Exemplo response (GET): <br>
    {
        "convention": "teste",
        "created_at": null,
        "external_id": null,
        "id": 2,
        "interest_correction": "teste",
        "name": "%TESTE",
        "series_id": 4514,
        "updated_at": null
    }<br>
    Exemplo response (PUT): <br>
    {
        "message": "Record created successfully"
    }
    <br>

</code>
<br>
<code> 
    URL: api/admin/v1/tables/<:table_name:>/crud/ <br>
    CRUD da tabela selecionada <br>
    Método: DELETE/POST <br>
    Exemplo response (POST): <br>
    {
        "message": "Record created successfully"
    }<br><br>
    Exemploy Payload (DELETE) <br>
    {
        "ids": ["1","2","3","4"]
    }
    <br>
    Exemplo response (DELETE): <br>
    {
        "message": "Record deleted successfully"
    }
    <br>
</code>

<code> 
    URL: api/admin/v1/tables/<:table_name:>/crud/?id= <br>
    CRUD da tabela selecionada <br>
    Método: DELETE <br>
    Exemplo response (DELETE): <br>
    {
	    "message": "Record deleted successfully"
    }<br>

</code>
