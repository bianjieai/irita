#!/usr/bin/make -f

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
PACKAGES_UNITTEST=$(shell go list ./... | grep -v '/simulation' | grep -v '/cli_test' | grep -v 'modules/wasm')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
BINDIR ?= $(GOPATH)/bin
LEDGER_ENABLED ?= true
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
NetworkType := $(shell if [ -z ${NetworkType} ]; then echo "mainnet"; else echo ${NetworkType}; fi)

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=irita \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=irita \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
		  -X github.com/tendermint/tendermint/crypto/algo.Algo=sm2 \
		  -X github.com/bianjieai/irita/address.Bech32ChainPrefix=i \
		  -X github.com/bianjieai/irita/address.PrefixAcc=a \
		  -X github.com/bianjieai/irita/address.PrefixAddress=a \
		  -X github.com/tharsis/ethermint/types.EvmChainID=1223

buildflags = -X github.com/tendermint/tendermint/crypto/algo.Algo=sm2

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

# The below include contains the tools target.

all: tools install lint

# The below include contains the tools.
include contrib/devtools/Makefile

build: go.sum
ifeq ($(OS),Windows_NT)
	go build $(BUILD_FLAGS) -o build/irita.exe ./cmd/irita
else
	go build $(BUILD_FLAGS) -o build/irita ./cmd/irita
endif

build-linux: go.sum update-swagger-docs
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

build-contract-tests-hooks:
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/contract_tests.exe ./cmd/contract_tests
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/contract_tests ./cmd/contract_tests
endif

install: go.sum
	go install $(BUILD_FLAGS) ./cmd/irita

update-swagger-docs: statik proto-swagger-gen
	$(BINDIR)/statik -src=lite/swagger-ui -dest=lite -f -m
	@if [ -n "$(git status --porcelain)" ]; then \
        echo "\033[91mSwagger docs are out of sync!!!\033[0m";\
        exit 1;\
    else \
    	echo "\033[92mSwagger docs are in sync\033[0m";\
    fi
.PHONY: update-swagger-docs

########################################
### Tools & dependencies

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

draw-deps:
	@# requires brew install graphviz or apt-get install graphviz
	go get github.com/RobotsAndPencils/goviz
	@goviz -i ./cmd/irita -d 2 | dot -Tpng -o dependency-graph.png

clean:
	rm -rf snapcraft-local.yaml build/ tmp-swagger-gen/

distclean: clean
	rm -rf vendor/

proto-all: proto-tools proto-gen proto-swagger-gen

proto-gen:
	@./scripts/protocgen.sh

proto-swagger-gen:
	@./scripts/protoc-swagger-gen.sh

########################################
### Testing


test: test-unit
test-all: test-race test-cover

test-unit:
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ${PACKAGES_UNITTEST}

test-race:
	@VERSION=$(VERSION) go test -mod=readonly -race -tags='ledger test_ledger_mock' ./...

test-cover:
	@go test -mod=readonly -timeout 30m -race -coverprofile=coverage.txt -covermode=atomic -tags='ledger test_ledger_mock' ./...

lint:
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/statik/statik.go" | xargs gofmt -d -s
	go mod verify

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/*/statik.go" -not -path "*.pb.go" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/*/statik.go" -not -path "*.pb.go" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/*/statik.go" -not -path "*.pb.go" | xargs goimports -w -local github.com/bianjieai/irita

benchmark:
	@go test -mod=readonly -bench=. ./...


########################################
### Local validator nodes using docker and docker-compose
build-docker-iritanode:
	docker build -t bianjieai/irita .

localnet-init:
	@if ! [ -f build/nodecluster/node0/irita/config/genesis.json ]; then docker run --rm -v $(CURDIR)/build:/home bianjieai/irita irita testnet --v 4 --output-dir /home/nodecluster --chain-id irita-test --keyring-backend test --starting-ip-address 192.168.10.2 ; fi
	@echo "To install jq command, please refer to this page: https://stedolan.github.io/jq/download/"
	@cat build/nodecluster/node0/irita/config/genesis.json | jq '.app_state.auth.accounts+= [{"@type": "/cosmos.auth.v1beta1.BaseAccount","address":"iaa15qgqfqk8uuej8ykjcyf7nse5n2avph0m92cu4e"}]' | jq '.app_state.bank.balances+= [{"address":"iaa15qgqfqk8uuej8ykjcyf7nse5n2avph0m92cu4e","coins":[{"denom":"point","amount":"1000000000000"}]}]' > build/genesis_temp.json
	@sudo cp build/genesis_temp.json build/nodecluster/node0/irita/config/genesis.json
	@sudo cp build/genesis_temp.json build/nodecluster/node1/irita/config/genesis.json
	@sudo cp build/genesis_temp.json build/nodecluster/node2/irita/config/genesis.json
	@sudo cp build/genesis_temp.json build/nodecluster/node3/irita/config/genesis.json
	@rm build/genesis_temp.json
	@echo "Faucet address: iaa15qgqfqk8uuej8ykjcyf7nse5n2avph0m92cu4e"
	@echo "Faucet coin amount: 1000000000000point"
	@echo "Faucet key seed: tube lonely pause spring gym veteran know want grid tired taxi such same mesh charge orient bracket ozone concert once good quick dry boss"

localnet-start: localnet-init localnet-stop
	docker-compose up -d

localnet-stop:
	docker-compose down