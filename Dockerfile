FROM golang:latest AS builder
ENV GOPROXY "https://goproxy.cn,direct"
WORKDIR /go/src/github.com/wukongcloud/xxdns
COPY . /go/src/github.com/wukongcloud/xxdns/
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./user-service main.go

FROM alpine:latest
#RUN apk add --no-cache curl \
#    && adduser -D -h /app -u 1000 app
RUN adduser -D -h /app -u 1000 app
WORKDIR /app
COPY --from=builder /go/src/github.com/wukongcloud/xxdns/xxdns /app/xxdns
RUN chmod +x /app/user-service
COPY --from=builder /go/src/github.com/wukongcloud/xxdns/config ./config
EXPOSE 8080
USER 1000
CMD ["/app/xxdns"]