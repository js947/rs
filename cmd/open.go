package cmd

import (
	"github.com/spf13/cobra"
	"log"
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
