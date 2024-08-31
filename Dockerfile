FROM golang:1.23.0-alpine3.20 AS base

WORKDIR /app

COPY . /app/

RUN go build -o redmineticketapi ./cmd/web

FROM alpine:latest AS latest

COPY --from=base /app/redmineticketapi /redmineticketapi
COPY --from=base /app/ui /ui

EXPOSE 4000

RUN chmod +x /redmineticketapi

CMD [ "/redmineticketapi" ]