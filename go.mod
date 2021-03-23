module github.com/bianjieai/irita

go 1.15

require (
	github.com/99designs/keyring v1.1.6
	github.com/CosmWasm/wasmd v0.14.1-0.20210111145259-1acda43e5322
	github.com/bianjieai/iritamod v0.0.0-20210322112104-d973ecd9584a
	github.com/cosmos/cosmos-sdk v0.41.0
	github.com/cosmos/go-bip39 v1.0.0
	github.com/dvsekhvalnov/jose2go v0.0.0-20201001154944-b09cfaf05951
	github.com/ghodss/yaml v1.0.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/irisnet/irismod v1.3.2-0.20210220095847-3f198a478395
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mtibben/percent v0.2.1
	github.com/olebedev/config v0.0.0-20190528211619-364964f3a8e4
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/snikch/goodman v0.0.0-20171125024755-10e37e294daa
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/tendermint v0.34.3
	github.com/tendermint/tm-db v0.6.3
	github.com/tidwall/gjson v1.6.1 // indirect
	google.golang.org/genproto v0.0.0-20210114201628-6edceaf6022f
	google.golang.org/grpc v1.35.0
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.34.4-0.20210127105926-4ac7b5a35238
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
