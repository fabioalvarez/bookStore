# Build Stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
WORKDIR /go/src/app/cmd
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./main.go

# Final Stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=goapp Version=0.0.1
EXPOSE 3000


