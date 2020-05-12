# First stage
FROM golang:alpine as builder

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git

CMD mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/proxy-server/main.go

# Second stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

CMD ["source .env && ./main"]
