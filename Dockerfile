FROM golang:1.16-alpine

WORKDIR /app

# cai thu vien
COPY * ./
RUN go mod download

# khoi dong app
CMD go run main.go