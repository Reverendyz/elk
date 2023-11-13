FROM golang:1.21-alpine3.18

RUN apk add bash

WORKDIR /api

COPY . .

CMD [ "bash", "-c", "/api/entrypoint.sh" ]