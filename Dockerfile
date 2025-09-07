ARG GO_VERSION=1.25
FROM golang:${GO_VERSION}-bookworm as builder

LABEL maintainer="Jason Cameron katib@jasoncameron.dev"
LABEL org.opencontainers.image.source="https://github.com/Jasonlovesdoggo/katib"
LABEL description="A tool to get your most recent GitHub contributions"

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /katib .


FROM debian:bookworm
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=builder /katib /usr/local/bin/
CMD ["katib"]
