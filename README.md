# api-calculator

an api that performs the following operations:

### add
```http
http://localhost:8080/{value1},{value2}
```
### substract
```http
http://localhost:8080/{value1},{value2}
```
### divide
```http
http://localhost:8080/{value1},{value2}
```
### random
```http
http://localhost:8080/random/{value}
```
## Other endpoints

### metrics
```http
http://localhost:8080/metrics
```
### liveness
```http
http://localhost:8080/liveness
```
### readiness
```http
http://localhost:8080/readiness

```

## Usage as a server

```go run main.go```


## Usage as a service in a kubernetes cluster 

```kubectl port-forward svc/api-calculator -n 'namespace' 7171:80```

```http://localhost:7171/add/3,2```




