FROM golang:1.24.5 AS build
COPY . ./
RUN go build -o /bin/scnorion-worker .

FROM debian:latest
RUN apt-get update && apt install -y ca-certificates
COPY --from=build /bin/scnorion-worker /bin/scnorion-worker
WORKDIR /tmp
HEALTHCHECK --interval=30s --timeout=5s --start-period=30s --retries=3 \
  CMD /bin/scnorion-worker healthcheck || exit 1
ENTRYPOINT ["/bin/scnorion-worker"]