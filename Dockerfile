FROM golang:1.19.1 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/app/main.go

FROM alpine:3.16.2
RUN apk add findutils
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN addgroup -S llg_backend && \
        adduser -S llg_backend -G llg_backend
USER llg_backend
WORKDIR /app
COPY --chown=llg_backend:llg_backend --chmod=754 --from=builder /app/main .
COPY --chown=llg_backend:llg_backend --chmod=754 --from=builder /app/scripts/wait-for .
COPY --chown=llg_backend:llg_backend --chmod=754 scripts/start.sh .
ENTRYPOINT []
CMD ["./start.sh", "./main"]