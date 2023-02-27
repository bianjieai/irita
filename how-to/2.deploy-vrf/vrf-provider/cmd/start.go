package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/app/client"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/pkg/configs"
)

var (
	localConfig string

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start VRF provider.",
		Run: func(cmd *cobra.Command, args []string) {
			online()
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&localConfig, "CONFIG", "c", "", "config path: /opt/local.toml")
}

func online() {
	data, err := ioutil.ReadFile(localConfig)
	if err != nil {
		fmt.Println("Error: get config data error: ", err)
		return
	}
	config, err := readConfig(data)
	if err != nil {
		fmt.Println("Error: read config error: ", err)
		return
	}
	run(config)
}

func run(cfg *configs.Config) {
	client.Serve(cfg)
}

func readConfig(data []byte) (*configs.Config, error) {
	v := viper.New()
	v.SetConfigType("toml")
	reader := bytes.NewReader(data)
	err := v.ReadConfig(reader)
	if err != nil {
		return nil, err
	}
	cfg := configs.NewConfig()
	if err := v.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
