# Demo how to access k8s API in a Pod via go client

[offical docs](https://kubernetes.io/docs/tasks/run-application/access-api-from-pod/#accessing-the-api-from-within-a-pod)
<br>
[Go client library](https://github.com/kubernetes/client-go/)


## Init project
```
go mod init main.go
go mod tidy  
```

## Deploy
This will push to my perosnal repository, replace to yours.
```
# edit tag in Makefile
make push

# edit image tag in deployment.yaml
kubectl apply -f deployment.yaml
```

## Test
```
There are 186 pods in the cluster
Found lr-nginx-859d8f47bb-ld4wm pod in app namespace
There are 186 pods in the cluster
```