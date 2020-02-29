package cmd

import (
	"log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(openCmd)
}

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "open current job in browser",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("command 'open' not implemented")
	},
}
