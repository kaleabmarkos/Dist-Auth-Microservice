FROM golang:1.24.2-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/auth-service
RUN go build -o main .

CMD ["./main"]
