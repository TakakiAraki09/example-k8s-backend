FROM golang:1.23
WORKDIR /app
COPY . .
ADD go.mod go.sum server.go ./
RUN go mod download
RUN go build -o server /app/server.go
CMD /app/server
