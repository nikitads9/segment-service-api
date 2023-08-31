FROM golang:1.21.0-alpine AS builder

COPY . /github.com/nikitads9/segment-service-api/

WORKDIR /github.com/nikitads9/segment-service-api/

RUN go mod download
RUN go build -o ./bin/segment_service cmd/server/segment_service.go
RUN chown -R root ./bin/segment_service

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/nikitads9/segment-service-api/bin .
COPY --from=builder /github.com/nikitads9/segment-service-api/config.yml .

#CMD ["./segment_service", "--config", "config.yml"]