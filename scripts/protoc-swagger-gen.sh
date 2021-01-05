#!/usr/bin/env bash

set -eo pipefail


IRISMOD_VERSION=v1.1.1-0.20201229063925-7d7dad20f951
IRITAMOD_VERSION=v0.0.0-20210105030814-46869d8f262a
SDK_VERSION=v0.34.4-0.20210105022052-24f2fc27b94f
WASMD_VERSION=v0.13.1-0.20201217131318-53bbf96e9e87
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
rm -rf ./proto
