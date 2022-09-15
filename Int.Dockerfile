FROM alpine:latest 

ENV PROJECT_NAME=rest-service

RUN mkdir -p /usr/bin/testdata
COPY bin/integration /usr/bin/integration
CMD [ "/usr/bin/integration -test.v" ]