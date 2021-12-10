# Dockerfile
FROM alpine

ENTRYPOINT ["/usr/bin/sensibo-timer-guard"]
COPY sensibo-timer-guard /usr/bin/sensibo-timer-guard
