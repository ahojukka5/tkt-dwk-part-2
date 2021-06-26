FROM golang as build

COPY . /build
WORKDIR /build

RUN CGO_ENABLED=0 go build -o read_timestamp -a -ldflags '-s' read_timestamp.go

FROM alpine:latest
COPY --from=build /build/read_timestamp .
ENTRYPOINT ["./read_timestamp"]
