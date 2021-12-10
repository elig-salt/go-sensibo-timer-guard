# Dockerfile
FROM alpine

ARG BINARY_NAME=sensibo-timer-guard

ENTRYPOINT ["/usr/bin/sensibo-timer-guard"]
COPY ${BINARY_NAME} /usr/bin/sensibo-timer-guard
