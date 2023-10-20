FROM golang:1.20 AS build-stage

WORKDIR /probeServer
COPY go.mod /probeServer/
RUN go mod download

COPY /src/*.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /latencyServer

FROM alpine:3.14 AS build-release-stage

WORKDIR /

COPY --from=build-stage /latencyServer /latencyServer
RUN apk add -U tzdata
ENV TZ=Europe/Sarajevo
RUN cp /usr/share/zoneinfo/Europe/Sarajevo /etc/localtime

ENTRYPOINT ["/latencyServer", "-port"]

