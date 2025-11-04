# shopping-cart-go

[![Go Build](https://github.com/conorheffron/shopping-cart-go/actions/workflows/go.yml/badge.svg)](https://github.com/conorheffron/shopping-cart-go/actions/workflows/go.yml)

[![Quality gate](https://sonarcloud.io/api/project_badges/quality_gate?project=conorheffron_shopping-cart-go)](https://sonarcloud.io/summary/new_code?id=conorheffron_shopping-cart-go)

[Sonar Scan & Test Coverage - Overall Summary](https://sonarcloud.io/summary/overall?id=conorheffron_shopping-cart-go&branch=main)

### Tech
 - Go 1.24, gocover-cobertura, SonarQube

### Build / Install
```
go build
```
### Run Tests 
```
go test -v
```
### Run API App
```
go run main.go
```
##### OR
```
./shopping-cart-go
```

### POST Request to Add Item 1
```
curl -X POST -H "Content-Type: application/json" -d '{"id":1,"name":"Laptop","price":999.99}' http://localhost:8080/add

{"id":1,"name":"Laptop","price":999.99}
```
### POST Request to Add Item 2
```
curl -X POST -H "Content-Type: application/json" -d '{"id":2,"name":"basketball","price":77.99}' http://localhost:8080/add

{"id":2,"name":"basketball","price":77.99}
```
### GET Items
```
curl http://localhost:8080/cart

[{"id":1,"name":"Laptop","price":999.99},{"id":2,"name":"basketball","price":77.99}]
```
