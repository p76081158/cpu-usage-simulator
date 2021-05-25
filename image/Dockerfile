FROM golang:1.14-alpine3.12

LABEL description="p76081158/cpu-usage-simulator"

RUN apk update && \
    apk add --no-cache stress-ng git && \
    git clone https://github.com/p76081158/cpu-usage-simulator.git && \
    cd cpu-usage-simulator && \
    go build && \
    rm -rf /var/cache/apk/*

WORKDIR /go/cpu-usage-simulator

# Default command
ENTRYPOINT [ "./cpu-usage-simulator" ]

# Exposed ports
EXPOSE 9090
