package cmd

import (
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "file [upload,list,cat,download,delete]",
}

func init() {
	rootCmd.AddCommand(fileCmd)
}
