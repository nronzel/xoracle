# syntax=docker/dockerfile:1

FROM debian:latest

ENV PORT 8080

COPY xoracle usr/bin/xoracle

CMD ["/bin/xoracle"]
