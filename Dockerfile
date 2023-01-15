FROM golang:1.19-alpine3.17

RUN apk add --no-cache git

WORKDIR /app/chonizabot

COPY go.mod .

RUN go mod download

COPY . .

ENTRYPOINT ["go", "run", "main.go"]