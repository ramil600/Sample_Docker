FROM golang:1.12-alpine
RUN apk add --no-cache git
ENV GO111MODULE=on
RUN mkdir -p /go/src/go-with-compose
WORKDIR /go/src/go-with-compose
COPY . .
RUN go get gopkg.in/redis.v4
CMD go run main.go