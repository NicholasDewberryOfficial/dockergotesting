FROM golang:1.22
WORKDIR /base
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o main main.go
EXPOSE 8080
cmd ["./main"]