FROM golang:1.23

ENV ROOT=/go/src/app
WORKDIR ${ROOT}
RUN apt update \
    && apt clean \
    && rm -r /var/lib/apt/lists/*

RUN apt install git && \
    apt install curl
COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080
CMD ["go", "run", "server.go"]
