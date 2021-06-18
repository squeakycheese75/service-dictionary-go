# service-dictionary-go

## service-dictionary-api

A Golang REST API built whilst attempting to keep to clean architecture principle.

## Dependancies



### Routing

Supports Mux & Chi

### Data

Supports SQLite, Postgres, ...

### Messaging

Supports RabbitMq, ...

### Authentiction

Supports JWT, ...

## Docker Support

```bash
docker build -t service-dictionary-api .
```

```bash
docker run -d --name myApp -p 8080:10000 service-dictionary-api
```
### Images
https://hub.docker.com/r/squeskycheese/service-dictionary-api


## Running

```bash
go run main.go   
```
## Testing

```bash
go test service/*.go -v
```

# Credits

https://gitlab.com/pragmaticreviews/golang-mux-api