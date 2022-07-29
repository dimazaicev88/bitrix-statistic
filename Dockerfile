FROM golang:1.18.4

RUN mkdir -p /var/server/

WORKDIR /var/server/

COPY . /var/server/
RUN go install github.com/go-delve/delve/cmd/dlv@v1.9.0
RUN go mod vendor
RUN go build -gcflags="all=-N -l" -o ./cmd/server ./cmd/server.go

EXPOSE 3000 40000

ENTRYPOINT ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "./cmd/server"]
#CMD ["./cmd/server"]
