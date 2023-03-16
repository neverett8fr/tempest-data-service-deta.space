# syntax=docker/dockerfile:1

## Build
FROM golang:1.19.2-buster AS builder

WORKDIR /tempest-data-service-deta.space

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY config/*.yaml ./

COPY . .
COPY *.go ./

RUN go build -o /tempest-data-service-deta.space

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /tempest-data-service-deta.space ./

EXPOSE 8080

ENTRYPOINT ["/tempest-data-service-deta.space"]