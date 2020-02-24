FROM golang:1.13.8-alpine3.11

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o reviewApi


ENTRYPOINT ["./reviewApi"] 

EXPOSE 8080
