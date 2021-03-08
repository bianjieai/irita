#
# Build image: docker build -t bianjie/irita .
#
FROM golang:1.15-alpine3.12 as builder

# this comes from standard alpine nightly file
#  https://github.com/rust-lang/docker-rust-nightly/blob/master/alpine3.12/Dockerfile
# with some changes to support CosmWasm toolchain, etc
RUN set -eux; apk add --no-cache ca-certificates build-base;

# Set up dependencies
ENV PACKAGES make gcc git libc-dev bash openssl

WORKDIR /irita

# Add source files
COPY . .

# Install minimum necessary dependencies
RUN apk add $PACKAGES

# NOTE: add these to run with LEDGER_ENABLED=true
# RUN apk add libusb-dev linux-headers

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v0.13.0/libwasmvm_muslc.a /lib/libwasmvm_muslc.a
RUN sha256sum /lib/libwasmvm_muslc.a | grep 39dc389cc6b556280cbeaebeda2b62cf884993137b83f90d1398ac47d09d3900

RUN LEDGER_ENABLED=false BUILD_TAGS=muslc make build

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