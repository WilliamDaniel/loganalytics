# WIP: Log Analytics
Realiza a leitura e processamento de logs gerados por API Gateway, exporta os dados processados em forma de relatório CSV. 
O sistema irá permitir carregar logs que estão no filesystem, mas também é flexível para implementar outras formas de leitura de logs. Como um adapter para ler logs armazenados no `Prometheus`, por exemplo. Para realizar todas as operações, o loganalytics conta com 4 serviços descritos a baixo. 
####  `Obs.: como o sistema não foi concluído, o código fonte ainda não representa todo o fluxo.`
## logreader
logreader é o serviço responsável por carregar cada linha do `json` de log dentro de um `slice` de string. 

## logparser
logparser é o serviço responsável por extrair os dados essenciais do log colocar em uma estrutura que possa ser utilizada pelo serviço de persistência de logs no
banco de dados. 
#### `Exemplo da estrutura preenchida`
```golang
parsedLog := ParsedLog{
		AuthenticatedEntity: RequestAuthenticatedEntity{
			ConsumerID: AuthenticatedEntityConsumerID{
				UUID: "b3467028-a118-3bbe-988f-35906729991c",
			},
		},
		Service: RequestService{
			ID: "22f8e3a6-01f7-3264-b4b5-9d178df11d06",
		},
		Latencies: RequestLatencies{
			Proxy: 1586,
			Gateway: 15,
			Request: 1882,
		},
	}
```

## logstorer
logstorer é o serviço responsável por persistir na base de dados os dados da estrutura que contém os dados já parseados. (`ParsedLog`)

## logexporter
logexporter é o serviço responsável por consutar os logos na base de dados e exportar relatórios em forma de .csv. 

### WIP: Como executar o sistema
O entrypoint da aplicação ainda não foi concluído e por isso se rodar o comando 
``go run ./...`` com o terminal na raíz do projeto, verá que o que ele faz é chamar todos os serviços passando suas implementações (adapters) e realizar a inserção dos logs presente no arquivo `logs.txt` dentro do banco de dados em memória. Foi deixado assim para fins de teste, mas ao finalizar a aplicação deverá ter um endpoint que receberá uma request com o payload: 

```json
Ex. 1
POST /process/log
{
    "command": "PROCESS_FROM_FILESYSTEM",
    "source":"backhole/logs.txt"
}
```
```json
Ex. 2
POST /process/log
{
    "command": "PROCESS_FROM_BUCKET",
    "source":"https://urlbucket.com/logsbucketinsomewhere"
}
```
```json
Ex. 3
POST /process/log
{
    "command": "PROCESS_FROM_QUEUE",
    "source":"kafka-topic"
}
```
### Testes
Para rodar os testes unitários, pode executar o comando `go test ./...`, na raíz do projeto, ou apenas dentro da pasta do serviço observado.

### Arquitetura
Tudo que faz sentido estar próximo ao serviço (entidade de erros, entidade de negócio, interfaces/gateways, etc.), foi incluído dentro do mesmo diretório. Contudo, houve a preocupação em manter a separação das camadas.


`Exemplo de log aceito como input no sistema`

```json
{
    "request": {
        "method": "GET",
        "uri": "/get",
        "url": "http://httpbin.org:8000/get",
        "size": "75",
        "querystring": {},
        "headers": {
            "accept": "*/*",
            "host": "httpbin.org",
            "user-agent": "curl/7.37.1"
        },
    },
    "upstream_uri": "/",
    "response": {
        "status": 200,
        "size": "434",
        "headers": {
            "Content-Length": "197",
            "via": "kong/0.3.0",
            "Connection": "close",
            "access-control-allow-credentials": "true",
            "Content-Type": "application/json",
            "server": "nginx",
            "access-control-allow-origin": "*"
        }
    },
    "authenticated_entity": {
        "consumer_id": "80f74eef-31b8-45d5-c525-ae532297ea8e"
    },
    "route": {
        "created_at": 1521555129,
        "hosts": null,
        "id": "75818c5f-202d-4b82-a553-6a46e7c9a19e",
        "methods": ["GET","POST","PUT","DELETE","PATCH","OPTIONS","HEAD"],
        "paths": [
            "/example-path"
        ],
        "preserve_host": false,
        "protocols": [
            "http",
            "https"
        ],
        "regex_priority": 0,
        "service": {
            "id": "0590139e-7481-466c-bcdf-929adcaaf804"
        },
        "strip_path": true,
        "updated_at": 1521555129
    },
    "service": {
        "connect_timeout": 60000,
        "created_at": 1521554518,
        "host": "example.com",
        "id": "0590139e-7481-466c-bcdf-929adcaaf804",
        "name": "myservice",
        "path": "/",
        "port": 80,
        "protocol": "http",
        "read_timeout": 60000,
        "retries": 5,
        "updated_at": 1521554518,
        "write_timeout": 60000
    },
    "latencies": {
        "proxy": 1430,
        "kong": 9,
        "request": 1921
    },
    "client_ip": "127.0.0.1",
    "started_at": 1433209822425
}
```
`latencies` contém alguns dados sobre as latências envolvidas:

-   `proxy`  é o tempo levado pelo serviço final para processar a requisição.
-   `kong`  é a latência referente a execução de todos os plugins pelo Kong (gateway).
-   `request`  é o tempo decorrido entre o primeiro byte ser lido do cliente e o último byte ser enviado a ele. Útil para detectar clientes lentos.
