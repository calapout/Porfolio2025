FROM golang:1.25.0-alpine as build
LABEL authors="calapout"

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go mod download
RUN go fmt
RUN mkdir bin
RUN go build -o ./bin/backend

FROM alpine:latest as prod
LABEL authors="calapout"

WORKDIR /app
COPY --from=build /app/bin/backend .

ENTRYPOINT ["/app/backend"]
