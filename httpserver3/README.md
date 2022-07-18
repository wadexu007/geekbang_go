## This is an execrise project about order pizzas

Write by Golang
Containerized by Docker
Deployed to Kubernetes [[Manifests](https://github.com/wadexu007/geekbang_go/tree/main/httpserver3/k8s-mainfest)] 

## How to run locally
```
go run main.go
or  
make local
```


## healthz check
```
curl -X GET 'http://localhost/healthz'
```

## Insert pizzas data
```
curl -X POST 'http://localhost/pizzas' -d '{"id":1,"name":"Pepperoni","price":12}' | jq
curl -X POST 'http://localhost/pizzas' -d '{"id":2,"name":"Capricciosa","price":10}' | jq
curl -X POST 'http://localhost/pizzas' -d '{"id":3,"name":"Margherita","price":15}' | jq
```

## Insert orders data
```
curl -X POST 'http://localhost/orders' -d '{"pizza_id":1,"quantity":3}' | jq
curl -X POST 'http://localhost/orders' -d '{"pizza_id":2,"quantity":2}' | jq
```

## Query data
```
curl -X GET 'http://localhost/pizzas'
curl -X GET 'http://localhost/orders'
curl -X GET 'http://localhost/orders/1' | jq 
curl -X GET 'http://localhost/orders/2' | jq
```
