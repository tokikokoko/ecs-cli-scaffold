FROM golang:1.15-alpine as builder

COPY main.go go.mod /http-server/
WORKDIR /http-server
RUN go build

FROM alpine:3
COPY --from=builder /http-server/http-server /bin/
EXPOSE 8080
ENTRYPOINT ["http-server"]
