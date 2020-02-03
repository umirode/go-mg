FROM golang:1.13.6-alpine AS build

WORKDIR /go/src/app

RUN apk add --no-cache git make bash protobuf && \
    go get github.com/golang/protobuf/protoc-gen-go

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN \
    make install-modules && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /go/bin/app


FROM scratch

COPY --from=build /go/bin/app /bin/app

ENTRYPOINT ["/bin/app"]