# Log Analytics
Realiza a leitura e processamento de logs gerados por API Gateway, exporta os dados processados em forma de relatório CSV

## FileProcessor
FileProcessor é o serviço responsável por processar os dados extraídos dos logs.

`Exemplo de log aceito como input no sistema`

```
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
