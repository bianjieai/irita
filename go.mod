module github.com/bianjieai/irita

go 1.15

require (
	github.com/99designs/keyring v1.1.6
	github.com/CosmWasm/wasmd v0.13.1-0.20201217131318-53bbf96e9e87
	github.com/bianjieai/iritamod v0.0.0-20210105030814-46869d8f262a
	github.com/cosmos/cosmos-sdk v0.40.0-rc5
	github.com/cosmos/go-bip39 v1.0.0
	github.com/dvsekhvalnov/jose2go v0.0.0-20201001154944-b09cfaf05951
	github.com/ghodss/yaml v1.0.0
	github.com/gorilla/mux v1.8.0
	github.com/irisnet/irismod v1.1.1-0.20201229063925-7d7dad20f951
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
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/tendermint v0.34.0
	github.com/tendermint/tm-db v0.6.3
	github.com/tidwall/gjson v1.6.1 // indirect
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.34.4-0.20210105022052-24f2fc27b94f
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.0-irita-210104
)
