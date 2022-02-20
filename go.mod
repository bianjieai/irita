module github.com/bianjieai/irita

go 1.17

require (
	github.com/99designs/keyring v1.1.6
	github.com/CosmWasm/wasmd v0.0.0-00010101000000-000000000000
	github.com/bianjieai/iritamod v1.2.0
	github.com/bianjieai/tibc-go v0.2.0
	github.com/cosmos/cosmos-sdk v0.44.4
	github.com/cosmos/go-bip39 v1.0.0
	github.com/dvsekhvalnov/jose2go v0.0.0-20201001154944-b09cfaf05951
	github.com/ethereum/go-ethereum v1.10.11
	github.com/ghodss/yaml v1.0.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/improbable-eng/grpc-web v0.15.0
	github.com/irisnet/irismod v1.5.2-0.20220220101314-1e155de674b5
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mtibben/percent v0.2.1
	github.com/olebedev/config v0.0.0-20190528211619-364964f3a8e4
	github.com/palantir/stacktrace v0.0.0-20161112013806-78658fd2d177
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/rs/cors v1.8.0
	github.com/snikch/goodman v0.0.0-20171125024755-10e37e294daa
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.9.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/tendermint v0.35.0
	github.com/tendermint/tm-db v0.6.4
	github.com/tharsis/ethermint v0.8.1
	golang.org/x/crypto v0.0.0-20211115234514-b4de73f9ece8
	google.golang.org/genproto v0.0.0-20211116182654-e63d96a377c4
	google.golang.org/grpc v1.42.0
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	//github.com/CosmWasm/wasmd => github.com/provenance-io/wasmd v0.19.0
	github.com/CosmWasm/wasmd => github.com/bianjieai/wasmd v0.19.1-0.20211215102105-45e28c7c896c
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.44.2-irita-20211102
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.8-irita-210413.0.20211012090339-cee6e09e8ae3
	github.com/tharsis/ethermint => github.com/bianjieai/ethermint v0.8.2-0.20220211020007-9ec25dde74d4
)
