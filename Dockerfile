FROM golang:alpine

RUN apk update && apk add git

WORKDIR /app

COPY . .

USER root

RUN chmod 644 /etc/letsencrypt/live/capstone.hanifz.com/fullchain.pem

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT [ "./binary" ]
