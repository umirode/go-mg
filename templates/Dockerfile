FROM golang:1.13.6-alpine AS build

WORKDIR /go/src/app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go get -d -v ./...

COPY . .
RUN \
    go generate && \
    make install-modules && \
    go build -ldflags "-s -w" -o /go/bin/app

FROM scratch

COPY --from=build /go/bin/app /bin/app

ENTRYPOINT ["/bin/app"]