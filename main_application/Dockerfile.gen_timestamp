FROM golang as build

COPY . /build
WORKDIR /build

RUN CGO_ENABLED=0 go build -o gen_timestamp -a -ldflags '-s' gen_timestamp.go

FROM scratch
COPY --from=build /build/gen_timestamp .
ENTRYPOINT ["./gen_timestamp"]
