package cmd

import (
	"os"
	"fmt"
	"log"
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
  
  func Execute() {
	rootCmd.PersistentFlags().StringP("token", "t", "", "Rescale API token")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))

	viper.SetConfigName(".rescale")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("RESCALE")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal(err)
		}
	}

	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }