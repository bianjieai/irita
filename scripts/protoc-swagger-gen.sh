#!/usr/bin/env bash

set -eo pipefail


IRISMOD_VERSION=v1.1.1-0.20210129120628-951dfea557b5
IRITAMOD_VERSION=v0.0.0-20210125115338-c52e28ce6cf7
SDK_VERSION=v0.34.4-0.20210127105926-4ac7b5a35238
WASMD_VERSION=v0.14.1-0.20210111145259-1acda43e5322
WASMD_PROTO_DIR=x/wasm/internal/types

chmod -R 755 ${GOPATH}/pkg/mod/github.com/irisnet/irismod@${IRISMOD_VERSION}/proto
chmod -R 755 ${GOPATH}/pkg/mod/github.com/bianjieai/iritamod@${IRITAMOD_VERSION}/proto
chmod -R 755 ${GOPATH}/pkg/mod/github.com/bianjieai/cosmos-sdk@${SDK_VERSION}/proto/cosmos
chmod -R 755 ${GOPATH}/pkg/mod/github.com/!cosm!wasm/wasmd@${WASMD_VERSION}/${WASMD_PROTO_DIR}

mkdir -p ./proto/wasm

cp -r ${GOPATH}/pkg/mod/github.com/irisnet/irismod@${IRISMOD_VERSION}/proto ./
cp -r ${GOPATH}/pkg/mod/github.com/bianjieai/iritamod@${IRITAMOD_VERSION}/proto ./
cp -r ${GOPATH}/pkg/mod/github.com/bianjieai/cosmos-sdk@${SDK_VERSION}/proto/cosmos ./proto
cp -r ${GOPATH}/pkg/mod/github.com/!cosm!wasm/wasmd@${WASMD_VERSION}/${WASMD_PROTO_DIR}/*.proto ./proto/wasm

sed -i "" "s@${WASMD_PROTO_DIR}@wasm@g" `grep -rl "${WASMD_PROTO_DIR}" ./proto/wasm`
mkdir -p ./tmp-swagger-gen

proto_dirs=$(find ./proto -path './proto/cosmos/base/tendermint*' -prune -o -name 'query.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do

    # generate swagger files (filter query files)
    query_file=$(find "${dir}" -maxdepth 1 -name 'query.proto')
    echo $query_file
    if [[ ! -z "$query_file" ]]; then
        protoc \
            -I "proto" \
            -I "third_party/proto" \
            "$query_file" \
            --swagger_out ./tmp-swagger-gen \
            --swagger_opt logtostderr=true --swagger_opt fqn_for_swagger_name=true --swagger_opt simple_operation_ids=true
    fi
done

# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine ./lite/config.json -o ./lite/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# clean swagger files
rm -rf ./tmp-swagger-gen

# clean proto files
rm -rf ./proto/cosmos

rm -rf ./proto/coinswap
rm -rf ./proto/htlc
rm -rf ./proto/nft
rm -rf ./proto/oracle
rm -rf ./proto/random
rm -rf ./proto/record
rm -rf ./proto/service
rm -rf ./proto/token

rm -rf ./proto/perm
rm -rf ./proto/identity
rm -rf ./proto/params
rm -rf ./proto/slashing
rm -rf ./proto/node
rm -rf ./proto/genutil
rm -rf ./proto/wasm
rm -rf ./proto/upgrade
