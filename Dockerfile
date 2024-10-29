FROM golang:alpine

RUN apk update && apk add chromium

WORKDIR /app

COPY . .

RUN go mod tidy 

RUN go build -o binary ./cmd/app

ENTRYPOINT ["/app/binary"]