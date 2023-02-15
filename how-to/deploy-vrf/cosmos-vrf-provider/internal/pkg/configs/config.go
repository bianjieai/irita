package configs

type (
	Config struct {
		App App `mapstructure:"app"`
		Eth Eth `mapstructure:"chain"`

		ContractServices ContractServices `mapstructure:"contract_services"`
	}

	ContractServices struct {
		VRF ServiceCfg `mapstructure:"vrf"`
	}

	ServiceCfg struct {
		ServiceName string      `mapstructure:"service_name"`
		Cache       Cache       `mapstructure:"cache"`
		Enabled     bool        `mapstructure:"enabled"`
		Contract    ContractCfg `mapstructure:"contract"`
	}
	// Eth config============================================================
	Eth struct {
		URI         string `mapstructure:"uri"`
		ChainID     uint64 `mapstructure:"chain_id"`
		GasLimit    uint64 `mapstructure:"gas_limit"`
		MaxGasPrice uint64 `mapstructure:"max_gas_price"`
		VrfAdminKey string `mapstructure:"vrf_admin_key"`
	}

	ContractCfg struct {
		Addr       string `mapstructure:"addr"`
		Topic      string `mapstructure:"topic"`
		OptPrivKey string `mapstructure:"opt_priv_key"`
	}

	Cache struct {
		Filename    string `mapstructure:"filename"`
		StartHeight uint64 `mapstructure:"start_height"`
	}

	App struct {
		Env      string `mapstructure:"env"`
		LogLevel string `mapstructure:"log_level"`
	}
)

func NewConfig() *Config {
	return &Config{}
}
