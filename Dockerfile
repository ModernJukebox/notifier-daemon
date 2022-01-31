##
## Build Stage
##
FROM golang:1.17 AS build

ENV GO111MODULE=on

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN GO111MODULE=on CGO_ENABLED=0 go build -o /go/bin/notifier-daemon

##
## Run Stage
##
FROM gcr.io/distroless/static

COPY --from=build /go/bin/notifier-daemon /

ENTRYPOINT ["/notifier-daemon"]