# hub.docker.com/Soul-Mate/webapp:latest

FROM golang:1.13.0-alpine3.10

COPY ./*.go /http_service/

EXPOSE 8888

CMD [ "go", "run", "/http_service/http_server.go", "/http_service/rpc_common.go"]
