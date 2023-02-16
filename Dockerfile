FROM golang:1.19 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o quicktmp .

FROM alpine:3.13 AS certer

RUN apk add -U --no-cache ca-certificates

FROM scratch

COPY --from=certer /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder build/quicktmp /usr/bin/quicktmp

CMD ["quicktmp", "serve"]