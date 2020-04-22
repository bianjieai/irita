module github.com/bianjieai/irita

require (
	github.com/btcsuite/btcd v0.20.1-beta // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/confio/go-cosmwasm v0.6.3
	github.com/cosmos/cosmos-sdk v0.38.2
	github.com/gorilla/mux v1.7.4
	github.com/irismod/nft v0.0.0-20200422071319-fd35c26b8173
	github.com/irismod/record v0.0.0-20200417015603-6b7b3ac5f2af
	github.com/irismod/service v0.0.0-20200422072147-1af06a4143ca
	github.com/irismod/token v0.0.0-20200422015846-1cf537b8f221
	github.com/keybase/go-keychain v0.0.0-20191114153608-ccd67945d59e // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/otiai10/copy v1.0.2
	github.com/otiai10/curr v0.0.0-20190513014714-f5a3d24e5776 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.2.1 // indirect
	github.com/prometheus/client_model v0.0.0-20191202183732-d1d2010b5bee // indirect
	github.com/prometheus/procfs v0.0.8 // indirect
	github.com/rakyll/statik v0.1.6
	github.com/rcrowley/go-metrics v0.0.0-20190826022208-cac0b30c2563 // indirect
	github.com/snikch/goodman v0.0.0-20171125024755-10e37e294daa
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v0.0.7
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.2
	github.com/tendermint/tm-db v0.5.1
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20191216205247-b31c10ee225f // indirect
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.38.2-0.20200422083600-e15718d623f2
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.0-irita-200302
)

go 1.13
