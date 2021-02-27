FROM golang:1.14

WORKDIR /go/src/github.com/Kirito055/Task1/client

COPY . /go/src/github.com/Kirito055/Task1/client

RUN apt-get update --yes
RUN apt-get install --yes netcat
RUN go mod download

EXPOSE 5555/tcp
