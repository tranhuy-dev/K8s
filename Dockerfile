FROM golang:1.13-alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
#download dependencies
RUN go mod download
COPY . .
RUN go build -o /go/bin/main
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./dist/main
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o ./dist/main
FROM alpine:3.5
RUN apk add --update ca-certificates
RUN apk add --no-cache tzdata && \
  cp -f /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
  apk del tzdata
WORKDIR /app
COPY --from=builder /go/bin/main .
#expose port
EXPOSE 9090
ENTRYPOINT ["./main"]