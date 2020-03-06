package cmd

import (
	"github.com/spf13/cobra"
)

var analysisCmd = &cobra.Command{
	Use:   "analysis",
	Short: "analysis [search,versions]",
}

func init() {
	rootCmd.AddCommand(analysisCmd)
}
