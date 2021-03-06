FROM golang:1.16-alpine as builder
RUN apk update && \
    apk add --virtual build-dependencies build-base
WORKDIR /srv
COPY . .
RUN make build

FROM alpine:latest
WORKDIR /srv
COPY --chown=0:0 --from=builder /srv/dist/ ./
EXPOSE 8080
CMD ["./go-pictonary"]
