#
# Build mainnet image: docker build -t irita/irita .
# Build testnet image: docker build -t irita/irita --build-arg NetworkType=testnet .
#
FROM golang:1.13-alpine3.10 as builder

# Set up dependencies
ENV PACKAGES make gcc git libc-dev bash linux-headers eudev-dev

WORKDIR /irita

# Add source files
COPY . .

# Install minimum necessary dependencies, run unit tests
RUN apk add --no-cache $PACKAGES && make test-unit

# Initialize network type, could be override via docker build argument `--build-arg NetworkType=testnet`
ARG NetworkType=mainnet

RUN make statik && make build

# ----------------------------

FROM alpine:3.10

# p2p port
EXPOSE 26656
# rpc port
EXPOSE 26657
# metrics port
EXPOSE 26660

COPY --from=builder /irita/build/ /usr/local/bin/