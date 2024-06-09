FROM golang:1.20.5-alpine AS build
WORKDIR /build
COPY . .
RUN go mod download
RUN go build

FROM alpine:latest as run
WORKDIR /bot
COPY --from=build /build/ .
RUN echo "nameserver 1.1.1.1" >> /etc/resolv.conf
CMD ./zenitria-bot
