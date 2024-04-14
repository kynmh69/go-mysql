FROM golang:1.22.2-bullseye as builder
WORKDIR /go/app
COPY ./ ./

RUN go build -a -o bin/main ./main.go

FROM alpine:latest
WORKDIR /go/app
COPY --from=builder /go/app/bin/main ./bin/
# CMD [ "bin/main" ]
