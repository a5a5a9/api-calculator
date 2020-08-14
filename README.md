# api-calculator

an api that performs the following operations:

### add
```http
http://localhost:8080/{value1},{value2} HTTP/1.1
```
### substract
```http
http://localhost:8080/{value1},{value2} HTTP/1.1
```
### divide
```http
http://localhost:8080/{value1},{value2} HTTP/1.1
```
### random
```http
http://localhost:8080/random/{value} HTTP/1.1
```
## Other endpoints

### metrics
```http
http://localhost:8080/metrics HTTP/1.1
```
### liveness
```http
http://localhost:8080/liveness HTTP/1.1
```
### readiness
```http
http://localhost:8080/readiness HTTP/1.1



```

## Usage as a server

```go run main.go``


## Usage as a service in a kubernetes cluster 

```kubectl port-forward svc/api-calculator -n 'namespace' 7171:80```
```http://localhost:7171/add/3,2``




