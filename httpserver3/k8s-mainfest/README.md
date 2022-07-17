This is an execrise project about order pizzas

# healthz check
```
curl -X GET 'https://exercise.audiencecuration.ai/healthz'
```

# insert pizzas data
```
curl -X POST 'https://exercise.audiencecuration.ai/pizzas' -d '{"id":1,"name":"Pepperoni","price":12}' | jq
curl -X POST 'https://exercise.audiencecuration.ai/pizzas' -d '{"id":2,"name":"Capricciosa","price":10}' | jq
curl -X POST 'https://exercise.audiencecuration.ai/pizzas' -d '{"id":3,"name":"Margherita","price":15}' | jq
```

# insert orders data
```
curl -X POST 'https://exercise.audiencecuration.ai/orders' -d '{"pizza_id":1,"quantity":3}' | jq
curl -X POST 'https://exercise.audiencecuration.ai/orders' -d '{"pizza_id":2,"quantity":2}' | jq
```

# query data
```
curl -X GET 'https://exercise.audiencecuration.ai/pizzas'
curl -X GET 'https://exercise.audiencecuration.ai/orders'
curl -X GET 'https://exercise.audiencecuration.ai/orders/1' | jq 
curl -X GET 'https://exercise.audiencecuration.ai/orders/2' | jq
```