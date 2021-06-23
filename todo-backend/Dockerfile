FROM golang as build

COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -a -ldflags '-s' .

FROM alpine:latest
COPY --from=build /app/todo-backend /app/todo-backend
EXPOSE 8000
ENTRYPOINT ["/app/todo-backend"]
