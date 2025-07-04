FROM golang:1.24.4-alpine3.22 AS base
WORKDIR /go/src/app
RUN apk update
RUN apk add curl

COPY go.mod go.sum ./
RUN go mod download

COPY . .

#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main -ldflags="-s -w" -trimpath
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM base AS local
RUN go install github.com/air-verse/air@v1.61.5
CMD ["air"]

FROM base AS production
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.0/migrate.linux-amd64.tar.gz | tar xvz
RUN mv ./migrate /usr/bin/migrate
CMD ["/go/src/app/main"]
