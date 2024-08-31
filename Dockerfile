FROM golang:1.23.0-alpine3.20 AS base

WORKDIR /app

COPY . /app/

RUN go build -o redmineticketapi ./cmd/web

FROM scratch

COPY --from=base /app/redmineticketapi /redmineticketapi
COPY --from=base /app/ui /ui

EXPOSE 4000

CMD [ "/redmineticketapi" ]