FROM alpine:edge

ENV CONFIG=""
EXPOSE 9096

COPY build/influxdb-relay /root/influxdb-relay
ENTRYPOINT ["/root/influxdb-relay"]