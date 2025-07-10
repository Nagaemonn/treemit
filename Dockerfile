FROM golang:1-bullseye AS builder
WORKDIR /work
ARG CGO_ENABLED=0
COPY . .
RUN go build -o treemit ./cmd/main/

FROM alpine:latest
ARG VERSION=0.5.1
LABEL org.opencontainers.image.source=https://github.com/Nagaemonn/treemit \
    org.opencontainers.image.version=${VERSION} \
    org.opencontainers.image.title=treemit \
    org.opencontainers.image.description="Enhanced \"tree\" command with file and depth limits."

RUN adduser -D -h /workdir nonroot && mkdir -p /workdir
COPY --from=builder /work/treemit /opt/treemit/treemit
COPY --from=golang:1.12 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /workdir
USER nonroot
ENTRYPOINT [ "/opt/treemit/treemit" ]