FROM golang:1.20.0-alpine3.17
WORKDIR /opt/golang/
COPY . .
CMD ["go","run", "./examples/fullapp/cmd/main.go"]
EXPOSE 8888
