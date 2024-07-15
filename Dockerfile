LABEL maintainer="Jason Cameron katib@jasoncameron.dev"
LABEL org.opencontainers.image.source="https://github.com/Jasonlovesdoggo/katib"
LABEL description="A tool to get your most recent GitHub contributions"
ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /katib .


FROM debian:bookworm
RUN sudo apt-get install -y ca-certificates
COPY --from=builder /katib /usr/local/bin/
CMD ["katib"]
