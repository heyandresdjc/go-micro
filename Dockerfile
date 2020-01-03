FROM golang:1.8 AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/app/main main
RUN ls
EXPOSE 80
ENTRYPOINT ["./main"]