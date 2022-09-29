FROM golang:1.19.1 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/app/main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:3.16.2
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate /usr/local/bin
COPY --from=builder /app/wait-for .
RUN chmod +x wait-for
COPY migrations ./migrations
COPY migrate.sh .
RUN chmod +x migrate.sh
COPY start.sh .
RUN chmod +x start.sh
RUN apk add findutils
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ENTRYPOINT []
CMD ["./start.sh", "./main"]