FROM golang as build

COPY . /build
WORKDIR /build

RUN CGO_ENABLED=0 go build -a -ldflags '-s' .

FROM scratch
COPY --from=build /build/main_application .
ENTRYPOINT ["./main_application"]
