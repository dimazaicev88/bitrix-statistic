FROM golang:1.18.4

RUN mkdir -p /var/server/

WORKDIR /var/server/

COPY . /var/server/

RUN go build  -o ./cmd/server ./cmd/server.go

EXPOSE 3000

CMD ["./cmd/server"]
