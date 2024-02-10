FROM debian:stable-slim

ENV PORT 8080

COPY xoracle usr/bin/xoracle

CMD ["/bin/xoracle"]
