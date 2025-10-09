# shopping-cart-golang

[![Go Build](https://github.com/conorheffron/shopping-cart-golang/actions/workflows/go.yml/badge.svg)](https://github.com/conorheffron/shopping-cart-golang/actions/workflows/go.yml)

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
./shopping-cart 
```

### POST Request to Add Item 1
```
shopping-cart-golang % curl -X POST -H "Content-Type: application/json" -d '{"id":1,"name":"Laptop","price":999.99}' http://localhost:8080/add

{"id":1,"name":"Laptop","price":999.99}
```
### POST Request to Add Item 2
```
shopping-cart-golang % curl -X POST -H "Content-Type: application/json" -d '{"id":2,"name":"basketball","price":77.99}' http://localhost:8080/add

{"id":2,"name":"basketball","price":77.99}
```
### GET Items
```
shopping-cart-golang % curl http://localhost:8080/cart

[{"id":1,"name":"Laptop","price":999.99},{"id":2,"name":"basketball","price":77.99}]
```
