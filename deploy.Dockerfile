# syntax=docker/dockerfile:1
# Dev

# Step 1 - Build
FROM golang:alpine AS build

ENV CGO_ENABLED 0

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download


COPY . .

RUN go build -o application cmd/app/main.go

# Step 2 - start
FROM alpine:3.19.1

RUN apk add --no-cache tzdata

WORKDIR /app

COPY /external ./external
COPY .env.deploy .
COPY --from=build /app .


EXPOSE 7771
EXPOSE 7779


# Config
ENV PROJECT_ABS_PATH "/app"

ENTRYPOINT ["./application", "-mode=deploy"]
