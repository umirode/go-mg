FROM golang:1.13.6-alpine AS builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk add --no-cache git protoc ca-certificates tzdata openssh && \
    go get github.com/golang/protobuf/protoc-gen-go && \
    go get -u github.com/umirode/prot && \
    \
    adduser -s /bin/true -u 1000 -D -h /app app && \
    sed -i -r "/^(app|root)/!d" /etc/group /etc/passwd && \
    sed -i -r 's#^(.*):[^:]*$#\1:/sbin/nologin#' /etc/passwd

WORKDIR /app

COPY prot.yaml ./
RUN prot install

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s -extldflags "-static"' -o /app/bin .

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/group /etc/shadow /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER app

COPY --from=builder /app/bin /app

ENTRYPOINT ["/app"]