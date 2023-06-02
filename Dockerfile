FROM golang:alpine

RUN apk update && apk add git

# Install Nginx
RUN apk add --no-cache nginx

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

# Expose ports
EXPOSE 80
EXPOSE 443

# Run Nginx in the background
# CMD nginx -g "daemon off;"
ENTRYPOINT [ "./binary" ]
