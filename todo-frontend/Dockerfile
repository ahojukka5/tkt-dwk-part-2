FROM golang as build

COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -a -ldflags '-s' .

FROM alpine:latest
COPY --from=build /app/ /app/
EXPOSE 3000
WORKDIR /app
CMD ["./todo-frontend"]
