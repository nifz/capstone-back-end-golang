FROM golang:alpine

RUN apk update && apk add git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

USER root

RUN chmod 644 /etc/letsencrypt/live/capstone.hanifz.com/fullchain.pem

USER nobody

ENTRYPOINT ["./binary"]
