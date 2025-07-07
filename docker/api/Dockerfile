FROM golang:1.23.4-alpine3.21 AS base
WORKDIR /go/src/app
RUN apk update
RUN apk add curl

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -trimpath -o main .

FROM base AS local
RUN go install github.com/air-verse/air@v1.61.5

CMD ["air"]

FROM base AS production
FROM alpine:3.21

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=base /go/src/app/main .
EXPOSE 50051

CMD ["./main"]
