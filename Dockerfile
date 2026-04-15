FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o main ./main.go

FROM alpine:latest

COPY --from=builder /app/main /app/main

CMD ["./main"]

EXPOSE 8080