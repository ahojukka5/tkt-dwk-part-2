FROM golang as build

COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -a -ldflags '-s' .

FROM alpine:latest
COPY --from=build /app/ /app/
EXPOSE 8000
CMD ["/app/todo-backend"]
