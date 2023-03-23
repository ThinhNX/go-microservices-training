# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:alpine
WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
COPY . .
RUN GO111MODULE=on GOOS=linux GOARCH=amd64 go build cmd/main.go
EXPOSE 8080

CMD ["./main"]

# ... the rest of the Dockerfile is ...
# ...   omitted from this example   ...