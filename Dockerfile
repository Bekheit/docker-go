FROM golang:latest

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

RUN go mod download

COPY . /app

EXPOSE 8080

CMD ["go", "run", "main.go"]