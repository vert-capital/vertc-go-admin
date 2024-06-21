# Documentação

## Objetivo
Ser uma biblioteca que aplicações da VERT usando GORM poderão criar automaticamente rotas para o ADMIN

## Rotas
<code> 
    URL: api/admin/v1/tables <br>
    Listagem de tabelas <br>
    Método: GET <br>
    Exemplo response: <br>
</code>
<br>
<code> 
    URL: api/admin/v1/tables/<:table_name:> <br>
    Estrutura de dados da tabela selecionada <br>
    Método: GET <br>
    Exemplo response: <br>
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
    ?action=xx -> Chamando ACTION <br>
    <br>
    Método: GET <br>
    Exemplo response: <br>
    Exemplo response (quando chamada a action): <br>
</code>
<br>
<code> 
    URL: api/admin/v1/tables/<:table_name:>/crud/<:id:> <br>
    CRUD da tabela selecionada <br>
    Método: GET/POST/PUT <br>
    Exemplo response (GET): <br>
    Exemplo response (POST/PUT): <br>
</code>
<br>
<code> 
    URL: api/admin/v1/tables/<:table_name:>/crud/ <br>
    CRUD da tabela selecionada <br>
    Método: DELETE/POST <br>
    Exemplo response (POST): <br>
    Exemploy Payload (DELETE) <br>
    {
        "ids": [1,2,3,4]
    }
    <br>
    Exemplo response (DELETE): <br>
</code>