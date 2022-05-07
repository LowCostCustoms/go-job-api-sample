FROM golang:1.18.1-alpine3.15 AS build

WORKDIR /app
COPY . .
RUN apk add --no-cache make protoc git
RUN CGO_ENABLED=0 GOOS=linux make build

FROM alpine:3.15

ENV MIGRATIONS_LOCATION="file:///app/migrations/"

WORKDIR /app
COPY --from=build /app/build/api .
COPY --from=build /app/migrations ./migrations

CMD ["/app/api", "serve"]
