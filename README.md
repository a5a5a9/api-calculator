# api-calculator

an api that performs the following operations:

### add
```http
http://localhost:8080/add/value1,value2
```
### substract
```http
http://localhost:8080/substract/value1,value2
```
### divide
```http
http://localhost:8080/division/value1,value2
```
### random (value is optional)
```http
http://localhost:8080/random/value
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


## as a service in a kubernetes cluster 

```kubectl port-forward svc/api-calculator -n 'namespace' 7171:80```

```http://localhost:7171/add/3,2```


## as ingress service based on dns name assigned by the nginx ingress controller

