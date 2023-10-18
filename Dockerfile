FROM golang:1.21.3 AS builder

WORKDIR /opt/workspace

COPY . .

RUN go build ./src/solax-exporter/

FROM alpine

COPY --from=builder /opt/workspace/solax-exporter .
COPY entrypoint.sh .

ENTRYPOINT [ entrypoint.sh ]