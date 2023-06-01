#
# Build image: docker build -t bianjie/irita .
#
FROM bianjie/irita-builder:v3 as builder

WORKDIR /irita
# Add source files
COPY . .

RUN go mod tidy

RUN LEDGER_ENABLED=false BUILD_TAGS=muslc make build

# ----------------------------

FROM bianjie/irita-runner:ubuntu

# Set up dependencies
ENV PACKAGES make gcc perl wget

WORKDIR /

# p2p port
EXPOSE 26656
# rpc port
EXPOSE 26657
# metrics port
EXPOSE 26660

COPY --from=builder /irita/build/irita /usr/local/bin/