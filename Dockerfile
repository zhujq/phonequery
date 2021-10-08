FROM golang:1.13.5-alpine3.10 AS builder
RUN apk update && apk add git && go get github.com/devfeel/dotweb && go get github.com/zhujq/phonedata
copy . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /phonequery phonequery.go



FROM alpine:latest
WORKDIR /
COPY --from=builder /phonequery .
copy . .
RUN  chmod +x /phonequery  && chmod 777 /entrypoint.sh
ENTRYPOINT  /entrypoint.sh 

EXPOSE 8080
