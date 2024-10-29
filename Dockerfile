FROM golang:1.23.1-alpine3.20

RUN apk update && apk add chromium

WORKDIR /app

COPY . .

RUN go mod tidy 

RUN go build -o binary .

ENTRYPOINT ["/app/binary"]