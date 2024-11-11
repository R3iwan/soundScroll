FROM golang:1.20-alpine

WORKDIR /app

COPY wait-for-it.sh /wait-for-it.sh

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

EXPOSE 8080

CMD /wait-for-it.sh postgres:5432 -- ./main