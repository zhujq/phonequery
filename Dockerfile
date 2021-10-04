FROM golang:1.16-alpine3.13 as builder

WORKDIR $GOPATH/src/phonequery
COPY . .
COPY phonedata $GOPATH/src/phonequery/phonedata

RUN apk add --no-cache git && set -x && \
    go mod init && go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o /phonequery phonequery-main.go


FROM alpine:latest

WORKDIR /
RUN mkdir -p /phonedata/
COPY --from=builder /phonequery . 
ADD ./phonedata/phone.dat  /phonedata/phone.dat

RUN  chmod +x /phonequery  && chmod 777 /entrypoint.sh
ENTRYPOINT  /entrypoint.sh 

EXPOSE 8080
