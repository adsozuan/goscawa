# BUILD
FROM golang:1.19.2-bullseye as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -v -o server ./cmd/goscawa/ 

# RUN
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

RUN groupadd -r goscawa && useradd --no-log-init -r -g goscawa goscawa

COPY --from=builder /app/server /app/server
RUN chown -R goscawa:goscawa /app

USER goscawa:goscawa

WORKDIR /app
EXPOSE 8080

ENTRYPOINT ["./server"]
