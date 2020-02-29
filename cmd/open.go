package cmd

import (
	"log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "open",
		Short: "Open current job in browser",
		Run: func(cmd *cobra.Command, args []string) {
			log.Fatal("command 'open' not implemented")
		},
	})
}
