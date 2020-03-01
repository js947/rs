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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hi")
	},
}

func Execute() error {
	rootCmd.PersistentFlags().StringP("token", "t", "", "Rescale API token")
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

	return rootCmd.Execute()
}
