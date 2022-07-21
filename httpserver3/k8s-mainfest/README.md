
## What's in this folder?

This folder contains a collection of Kubernetes yaml used to deploy [exercise project](https://github.com/wadexu007/geekbang_go/tree/main/httpserver3) in kubernetes clusters. 

## Including K8S features
* High Availability with replicas=2
* Pod Anti Affinity (pods fall in different nodes)
* Node Affinity
* QoS, resources request and limit
* Liveness/Readiness Probe
* PersistentVolume and PersistentVolumeClaim
* ConfigMap
* Taints and Tolerations
* Shutdown Gracefully
* Rolling Update
* Secruity
   * TLS Secret (automatic sign from Let's Encrypt via cert-manager)
   * Ingress Nginx
     * Whitelist access
     * Force ssl redirect (http -> https)
   * Multiple namespaces to separate different resources
     * Ingress gateway in dmz namespace
     * Service and Deploy in exercise namespace
   * ExternalName Service (Ingress in dmz -> external svc in dmz -> svc in exercise)
   * Not Allow Privilege Escalation

## Deployment
```
# deploy cert-manager
kubectl apply -f cert-manager.yaml
kubectl apply -f certificate-exercise.yaml

# deploy Ingress Gateway
kubectl apply -f ingress-controller.yaml
kubectl apply -f ingress.yaml

# deploy execrise project
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

## Upgrade
Refer to [How to build and release](https://github.com/wadexu007/geekbang_go/tree/main/httpserver3#how-to-build-and-release)
```
# Then edit image tag xxx in deployment.yaml
image: asia.gcr.io/devops-apac-mgmt/exercise-pizza:xxx

# Then deploy new version
kubectl apply -f deployment.yaml
```

## Test
### Healthz check
```
curl -X GET 'https://exercise.audiencecuration.ai/exercise/healthz'
```

### Insert pizzas data
```
curl -X POST 'https://exercise.audiencecuration.ai/exercise/pizzas' -d '{"id":1,"name":"Pepperoni","price":12}' | jq
curl -X POST 'https://exercise.audiencecuration.ai/exercise/pizzas' -d '{"id":2,"name":"Capricciosa","price":10}' | jq
curl -X POST 'https://exercise.audiencecuration.ai/exercise/pizzas' -d '{"id":3,"name":"Margherita","price":15}' | jq
```

### Insert orders data
```
curl -X POST 'https://exercise.audiencecuration.ai/exercise/orders' -d '{"pizza_id":1,"quantity":3}' | jq
curl -X POST 'https://exercise.audiencecuration.ai/exercise/orders' -d '{"pizza_id":2,"quantity":2}' | jq
```

### Query data
```
curl -X GET 'https://exercise.audiencecuration.ai/exercise/pizzas'
curl -X GET 'https://exercise.audiencecuration.ai/exercise/orders'
curl -X GET 'https://exercise.audiencecuration.ai/exercise/orders/1' | jq 
curl -X GET 'https://exercise.audiencecuration.ai/exercise/orders/2' | jq
```

## Logging
[logutils](https://github.com/hashicorp/logutils) from Hashicorp
```
2022/07/18 18:18:19 [INFO] Red Configuration
2022/07/18 18:18:19 [INFO] Start http server
2022/07/18 18:18:24 [ERROR] Can't read pizzas data from csv
2022/07/18 18:18:34 [INFO] Write pizza record to csv
2022/07/18 18:18:39 [INFO] get all pizzas
```
