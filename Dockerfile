#first stage - builder
FROM golang:1.13-stretch as builder
COPY . /skeltun
WORKDIR /skeltun
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o skeltun
COPY /script/wait-for-it.sh /usr/wait-for-it.sh
RUN chmod +x /usr/wait-for-it.sh
ENTRYPOINT ["/usr/wait-for-it.sh", "pgsql:5432"]

#second stage
FROM alpine:latest
WORKDIR /root/
RUN apk add --no-cache tzdata
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /skeltun .

ENTRYPOINT ["./skeltun"]