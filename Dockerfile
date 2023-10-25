# Dockerfile
FROM golang:1.20.4
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8085
CMD ["./main"]
