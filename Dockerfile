#
# Build image: docker build -t bianjieai/irita .
#
FROM golang:1.14.4-buster as builder

# Set up dependencies
ENV PACKAGES make gcc git libc-dev bash openssl

WORKDIR /irita

# Add source files
COPY . .

# Install minimum necessary dependencies, run unit tests
RUN apt-get update && apt-get install $PACKAGES -y && make test-unit

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v0.12.0/libwasmvm_muslc.a /lib/libwasmvm_muslc.a
RUN sha256sum /lib/libwasmvm_muslc.a | grep 00ee24fefe094d919f5f83bf1b32948b1083245479dad8ccd5654c7204827765

RUN BUILD_TAGS=muslc make build

# ----------------------------

FROM ubuntu:16.04

# Set up dependencies
ENV PACKAGES make gcc perl wget

WORKDIR /

# Install openssl 3.0.0
RUN apt-get update && apt-get install $PACKAGES -y \
    && wget https://github.com/openssl/openssl/archive/openssl-3.0.0-alpha4.tar.gz \
    && tar -xzvf openssl-3.0.0-alpha4.tar.gz \
    && cd openssl-openssl-3.0.0-alpha4 && ./config \
    && make install \
    && cd ../ && rm -fr *openssl-3.0.0-alpha4* \
    && apt-get remove --purge $PACKAGES -y && apt-get autoremove -y

# p2p port
EXPOSE 26656
# rpc port
EXPOSE 26657
# metrics port
EXPOSE 26660

COPY --from=builder /irita/build/ /usr/local/bin/