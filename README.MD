## Projeto com o objetivo de estudar GraphQL utilizando golang e a lib gqlgen

### Conceito

- GraphQL é uma linguagem de consulta para APIs e uma runtime para atender a essas consultas com seus
  dados existentes;
- GraphQL fornece uma descricao completa e compreensível dos dados em sua API;
- GraphQL dá aos clientes o poder de pedir exatamente o que eles precisam e nada mais;

### Vantagens

- Único endpoint;
- Única request para pegar tudo o que precisa(no rest seria feito uma request para cada parte da tela);
- Server apresenta os recursos disponíveis;
- Client solicita somente a informacao necessária;
- Há possibilidade de realizar alteracoes/insercoes através de "mutations";
- Retorna json como response;
- Utiliza protocolo HTTP;

### Schema

- GraphQL Schema Language;
- Data types;
- Query;
- Mutations;
- Subscriptions(Conexão via websocket para atualizacao de informacoes);

### Funcionamento

- Queries realizam consultas e trazem os dados de acordo com o solicitado. Executadas em paralelo;
- Mutation realiza o processo create, update e delete. É executada em série.

### Inicializacao do projeto

- Execute o comando `'go mod init github.com/USERNAME/REPO_NAME'`
- Instale o gqlgen com os comandos:
  `'printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go'`
  `'go mod tidy'`
- Gere os modelos com o comando:
  `'go run github.com/99designs/gqlgen init'`
- Para gerar novos modelos utilizando um schema.graphqls diferente, altere esse arquivo com o modelo que quer,
  apague as útlimas funcões do arquivo schema.resolvers.go e execute o comando:
  `'go run github.com/99designs/gqlgen generate'`
- Para rodar o projeto basta rodar o comando:
  `'go run server.go'`

### Conceitos da geracao do projeto

- Os models da api ficam armazenados na pasta graph/model;
- O resolver.go é onde ficam as dependências que serão implementadas no schema.resolvers.go;
- O schema.resolvers.go é o arquivo onde iremos implementar as funcoes para trazer/inserir dados;
