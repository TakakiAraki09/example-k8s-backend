FROM golang:1.23
WORKDIR /app
ADD go.mod go.sum server.go ./
ADD config/local.env ./
RUN go mod download
RUN go build -o server /app/server.go
CMD /app/server
