FROM golang:1.17.2-alpine3.13 as build

WORKDIR /app

COPY ./go.mod ./

RUN go mod download

COPY . .

RUN go build -o app ./cmd

FROM alpine:3.7

COPY --from=build /app/app /usr/local/bin/stress-test

RUN chmod +x /usr/local/bin/stress-test

ENTRYPOINT ["/usr/local/bin/stress-test"]