package cmd

import (
	"log"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Submit job",
		Run: submit,
	}

	cmd.PersistentFlags().String("config", "rescale", "job config file")
	cmd.PersistentFlags().StringP("path", "p", ".", "path to job")

	cmd.Flags().StringP("core", "c", "", "core type")
	viper.BindPFlag("core", cmd.Flags().Lookup("core"))
	cmd.Flags().IntP("numcores", "n", 0, "number of cores")
	viper.BindPFlag("numcores", cmd.Flags().Lookup("numcores"))

	rootCmd.AddCommand(cmd)
}

func submit(cmd *cobra.Command, args []string) {
	name, err := cmd.Flags().GetString("config")
	if err != nil {
		log.Fatal(err)
	}
	viper.SetConfigName(name)

	path, err := cmd.Flags().GetString("path")
	if err != nil {
		log.Fatal(err)
	}
	viper.AddConfigPath(path)

	viper.SetDefault("core", "onyx")
	viper.SetDefault("numcores", 2)

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("core      %s\n", viper.GetString("core"))
	fmt.Printf("num cores %d\n", viper.GetInt("numcores"))

}