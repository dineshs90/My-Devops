FROM golang:1.24.4-alpine

WORKDIR /app

COPY go/go.mod go/go.sum ./

RUN go mod download

COPY go/. . 

RUN go build -v -o my-app .

CMD ["./my-app"]


