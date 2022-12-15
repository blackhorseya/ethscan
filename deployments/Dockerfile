FROM golang:alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY api ./api
COPY cmd ./cmd
COPY pkg ./pkg
COPY internal ./internal

ARG MAIN_FOLDER

RUN GOOS=linux GOARCH=amd64 go build -o server ${MAIN_FOLDER}

FROM alpine:3 AS final

LABEL maintainer.name="sean.zheng"
LABEL maintainer.email="blackhorseya@gmail.com"

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /src/server ./server

ENTRYPOINT ["./server", "-c", "/app/configs.yaml"]
