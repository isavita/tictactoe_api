# syntax=docker/dockerfile:1
## Build
FROM golang:1.20-buster AS build

WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o /tictactoe-api ./cmd/server/main.go

## Run
FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=build /tictactoe-api /tictactoe-api
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/tictactoe-api"]
