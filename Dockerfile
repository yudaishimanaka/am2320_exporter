FROM quay.io/prometheus/busybox:latest

COPY am2320-exporter /bin/am2320-exporter

ENTRYPOINT ["/bin/am2320-exporter"]
EXPOSE 9430
