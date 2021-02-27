FROM golang

WORKDIR /go/src/github.com/Kirito055/Task1/server

COPY . /go/src/github.com/Kirito055/Task1/server

RUN go mod download

EXPOSE 5555/tcp

CMD go run server.go
