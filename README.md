(base) conorheffron@Conors-MacBook-Pro shopping-cart-golang % curl -X POST -H "Content-Type: application/json" -d '{"id":1,"name":"Laptop","price":999.99}' http://localhost:8080/add

{"id":1,"name":"Laptop","price":999.99}
(base) conorheffron@Conors-MacBook-Pro shopping-cart-golang % curl http://localhost:8080/cart

[{"id":1,"name":"Laptop","price":999.99}]
(base) conorheffron@Conors-MacBook-Pro shopping-cart-golang % curl -X POST -H "Content-Type: application/json" -d '{"id":1,"name":"basketball","price":77.99}' http://localhost:8080/add

{"id":1,"name":"basketball","price":77.99}
(base) conorheffron@Conors-MacBook-Pro shopping-cart-golang % curl http://localhost:8080/cart

[{"id":1,"name":"Laptop","price":999.99},{"id":1,"name":"basketball","price":77.99}]
(base) conorheffron@Conors-MacBook-Pro shopping-cart-golang % curl http://localhost:8080/cart

[{"id":1,"name":"Laptop","price":999.99},{"id":1,"name":"basketball","price":77.99}]
(base) conorheffron@Conors-MacBook-Pro shopping-cart-golang % curl http://localhost:8080/cart

curl: (52) Empty reply from server
(base) conorheffron@Conors-MacBook-Pro shopping-cart-golang % curl http://localhost:8080/cart

null
(base) conorheffron@Conors-MacBook-Pro shopping-cart-golang % curl -X POST -H "Content-Type: application/json" -d '{"id":1,"name":"basketball","price":77.99}' http://localhost:8080/add

{"id":1,"name":"basketball","price":77.99}
(base) conorheffron@Conors-MacBook-Pro shopping-cart-golang % curl http://localhost:8080/cart

go run main.go