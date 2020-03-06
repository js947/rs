package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "rs",
	Short: "rs is a simple cli for the rescale scaleX platform",
	Long: `rs is a simple cli for the rescale scaleX platform made by js947. 
See http://github.com/js947/rs`,
}

func Execute() error {
	rootCmd.PersistentFlags().StringP("token", "t", "unset", "Rescale API token")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	rootCmd.PersistentFlags().String("api", "", "API address")
	viper.BindPFlag("api", rootCmd.PersistentFlags().Lookup("api"))

	viper.SetDefault("api", "https://platform.rescale.com/api/")
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.rescale")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("RESCALE")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	if viper.GetString("token") == "unset" {
		return fmt.Errorf("must set rescale token in ~/.rescale/config.yaml or in RESCALE_TOKEN")
	}

	return rootCmd.Execute()
}
