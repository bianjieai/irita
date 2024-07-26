#!/usr/bin/env bash

set -eo pipefail
cd proto

proto_dirs=$(find ./irita -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    # this regex checks if a proto file has its go_package set to cosmossdk.io/api/...
    # gogo proto files SHOULD ONLY be generated if this is false
    # we don't want gogo proto to run for proto files which are natively built for google.golang.org/protobuf
    if grep -q "option go_package" "$file"; then
      buf generate --template buf.gen.gogo.yaml $file
    fi
  done
done

# for dir in $proto_dirs; do
#   protoc \
#   -I "proto" \
#   -I "third_party/proto" \
#   --gocosmos_out=plugins=interfacetype+grpc,\
# Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:. \
#   $(find "${dir}" -maxdepth 1 -name '*.proto')

#   # command to generate gRPC gateway (*.pb.gw.go in respective modules) files
#   protoc \
#   -I "proto" \
#   -I "third_party/proto" \
#   --grpc-gateway_out=logtostderr=true:. \
#   $(find "${dir}" -maxdepth 1 -name '*.proto')

# done
cd ..

# move proto files to the right places
cp -r github.com/bianjieai/irita/* ./
rm -rf github.com

./scripts/protocgen-pulsar.sh
