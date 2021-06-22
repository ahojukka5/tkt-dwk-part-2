FROM golang as build

COPY . /build
WORKDIR /build

RUN CGO_ENABLED=0 go build -a -ldflags '-s' .

FROM alpine:latest
COPY --from=build /build/app .
ENTRYPOINT ["./app"]
