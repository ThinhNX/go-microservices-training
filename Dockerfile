# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:alpine as builder
WORKDIR /app

# Download necessary Go modules
ADD go.mod go.sum ./
RUN go mod download
COPY . .
RUN GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o customer cmd/main.go

FROM alpine
COPY *.pem .
COPY --from=builder /app/customer . 
EXPOSE 8080
CMD ["./customer"]
# ... the rest of the Dockerfile is ...
# ...   omitted from this example   ...