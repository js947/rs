package cmd

import (
	"github.com/spf13/cobra"
)

var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "job [list,rename,delete]",
}

func init() {
	rootCmd.AddCommand(jobCmd)
}