FROM golang:1.15-alpine3.12 AS build-env
RUN apk add build-base linux-headers
ARG VERSION
ENV BUILD_HOME=/BUILD
RUN mkdir -p ${BUILD_HOME}
WORKDIR ${BUILD_HOME}

ADD . ${BUILD_HOME}
RUN go mod download
RUN go build -o /bin/go-defi \
    --tags "json1" \
    -ldflags "-X main.Version=$VERSION"

FROM alpine:3.12
COPY --from=build-env /bin/go-defi /bin/go-defi
ENTRYPOINT  ["/bin/go-defi", "--config.path", "/config.yaml" ]