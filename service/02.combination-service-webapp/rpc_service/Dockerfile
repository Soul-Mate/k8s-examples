# hub.docker.com/Soul-Mate/webapp:latest

FROM golang:1.13.0-alpine3.10

COPY ./*.go /rpc_service/

EXPOSE 30001

CMD [ "go", "run", "/rpc_service/arith_server.go", "/rpc_service/rpc_common.go"]