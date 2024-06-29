## Descrição
Este repositório contem todo o código relacionado a API SmartCleaner.

Para mais informações sobre a arquiteture e recursos do projeto, visite: [SmartCleaner Wiki](https://github.com/rafaelsouzak2b/smart-cleaner-ui/wiki)

<br>

**Link para acesso do projeto**

Para acessar o projeto em produção visite: [http://production.eba-32pvrax3.us-east-1.elasticbeanstalk.com/](http://production-api.eba-ecm2h3cp.us-east-1.elasticbeanstalk.com/api/v1/cleaner)

## Como executar o projeto

### Defina as variáveis de ambiente
Renomeie o arquivo .env.example para .env e substituia as informações padrão por credenciais válidas.<br>
Abaixo temos o arquivo .env.example com informações fictícias.

```
PORT=666

POSTGRES_HOST=localhost
POSTGRES_USER=xxx
POSTGRES_PASSWORD=XX
AWS_REGION=XXXX
AWS_IMG_PROFILE_BUCKET=XXX
AWS_ACCESS_KEY_ID='XXX'
AWS_SECRET_ACCESS_KEY='XXX'
AWS_SESSION_TOKEN=XXXX
DEFAULT_TOKEN='xxx'
CLEANER_TOKEN='xxx'
```

### Rode os comandos para instalação e execução do projeto

```shell
// para instalação dos pacotes
go mod tidy

// para rodar em modo de desenvolvimento
go run main.go

// para compilar a aplicação
go mod vendor

```





