FROM --platform=$BUILDPLATFORM golang:1.20 AS build-stage
ARG TARGETOS
ARG TARGETARCH

WORKDIR /probeServer
COPY go.mod ./
RUN go mod download

COPY /src/*.go ./
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /latencyServer

FROM alpine:3.14 AS build-release-stage

WORKDIR /

COPY --from=build-stage /latencyServer /latencyServer
RUN apk add -U tzdata
ENV TZ=Europe/Sarajevo
RUN cp /usr/share/zoneinfo/Europe/Sarajevo /etc/localtime
ENV PORT=1501

ENTRYPOINT ["/bin/sh", "-c","./latencyServer -port $PORT"]

