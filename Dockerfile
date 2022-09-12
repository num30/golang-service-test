FROM alpine:3.13.5

COPY bin/api-service /usr/bin/api-service

CMD ["/usr/bin/api-service"]