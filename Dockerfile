FROM golang:1.21.3 AS builder

WORKDIR /opt/workspace

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o solax_exporter .

FROM alpine

WORKDIR /opt/workspace

COPY --from=builder /opt/workspace/solax_exporter .
COPY entrypoint.sh .

ENTRYPOINT [ entrypoint.sh ]