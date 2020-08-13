FROM golang:1.13.8 as build

ENV GO111MODULE=on

RUN mkdir -p /go/src/webservice
WORKDIR /go/src/webservice
COPY . /go/src/webservice

RUN go mod download


RUN CGO_ENABLED=0 go build -a -installsuffix cgo --ldflags "-s -w" -o /usr/bin/server 

#final stage

FROM alpine:3.7

COPY --from=build /usr/bin/server /root/

EXPOSE 8080
WORKDIR /root/

CMD ["./server"]

