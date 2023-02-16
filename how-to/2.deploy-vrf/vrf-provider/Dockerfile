# Build image: docker build -t relayers/fabric .
FROM golang:1.19-alpine3.16 as builder

# Set up dependencies
ENV PACKAGES make git libc-dev bash gcc
WORKDIR $GOPATH/src
COPY . .
# Install minimum necessary dependencies, build binary
RUN apk add --no-cache $PACKAGES && make install

FROM alpine:3.16
COPY --from=builder /go/bin/vrf-provider /usr/local/bin/vrf-provider

CMD ["vrf-provider"]