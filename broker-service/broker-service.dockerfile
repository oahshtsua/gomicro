FROM alpine:3.20

WORKDIR /app
COPY broker-app .

EXPOSE 5000

CMD ["./broker-app"]

