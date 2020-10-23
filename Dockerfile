FROM golang:alpine as build-env

RUN apk update && apk add git

WORKDIR /src

COPY . .

RUN go mod tidy
RUN go build -o goech echFundamental/app

FROM alpine
WORKDIR /app
COPY --from=build-env /src/goech /app

ENTRYPOINT ./goech
