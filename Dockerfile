FROM golang:1.8 AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build .
RUN ls
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/app/app app
COPY --from=builder /go/src/app/vars vars
RUN ls
EXPOSE 8080
ENTRYPOINT ["./app"]