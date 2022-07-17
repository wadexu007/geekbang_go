# insert pizzas data
curl -X POST 'http://localhost:8080/pizzas' -d '{"id":1,"name":"Pepperoni","price":12}' | jq
curl -X POST 'http://localhost:8080/pizzas' -d '{"id":2,"name":"Capricciosa","price":10}' | jq
curl -X POST 'http://localhost:8080/pizzas' -d '{"id":3,"name":"Margherita","price":15}' | jq

# insert orders data
curl -X POST 'http://localhost:8080/orders' -d '{"pizza_id":1,"quantity":3}' | jq
curl -X POST 'http://localhost:8080/orders' -d '{"pizza_id":2,"quantity":2}' | jq

# query data
curl -X GET 'http://localhost:8080/healthz'
curl -X GET 'http://localhost:8080/pizzas'
curl -X GET 'http://localhost:8080/orders'
curl -X GET 'http://localhost:8080/orders/1' | jq 
curl -X GET 'http://localhost:8080/orders/2' | jq 