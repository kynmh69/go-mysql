FROM golang:1.22.2-bullseye as build

COPY ./ /go/app
WORKDIR /go/app

RUN go build -o ./bin/

FROM alpine
COPY --from=build /go/app/bin/main /usr/local/
CMD [ "main" ]