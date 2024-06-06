package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/services/keystore"
)

var (
	password  string
	genKeyCmd = &cobra.Command{
		Use:   "genkey",
		Short: "genrater key for vrf-provider.",
		Run: func(cmd *cobra.Command, args []string) {
			genKey()
		},
	}
)

func init() {
	rootCmd.AddCommand(genKeyCmd)
	genKeyCmd.Flags().StringVarP(&password, "PASSWORD", "p", "12345678", "key password, default: 12345678")
}

func genKey() {
	keyStoreMaster := keystore.New()
	vrfkey, err := keyStoreMaster.VRF().Create()
	if err != nil {
		fmt.Println("Error: Gen Key : ", err)
		return
	}
	fmt.Println(password)
	keyJSON, err := keyStoreMaster.VRF().Export(vrfkey.ID(), password)
	if err != nil {
		fmt.Println("Error: Export Key : ", err)
		return
	}
	fmt.Println(string(keyJSON))

}
