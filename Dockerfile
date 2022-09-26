FROM golang:1.18.6-alpine3.16

WORKDIR /server

COPY go.mod ./

COPY go.sum ./

COPY .env ./

RUN go mod download

COPY . .

RUN go build -o server server.go

EXPOSE 8080
EXPOSE 3306

CMD [ "./server" ]






