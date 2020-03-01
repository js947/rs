package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	files := &cobra.Command{
		Use:   "files",
		Short: "List files (shorthand for 'file list')",
		Run: func(cmd *cobra.Command, args []string) {
			err := file_list(cmd)
			if err != nil {
				panic(err)
			}
		},
	}
	rootCmd.AddCommand(files)

	list := &cobra.Command{
		Use:   "list",
		Short: "List files",
		Run: func(cmd *cobra.Command, args []string) {
			err := file_list(cmd)
			if err != nil {
				panic(err)
			}
		},
	}
	fileCmd.AddCommand(list)
}

func file_list(cmd *cobra.Command) error {
	return fmt.Errorf("file_list not implemented")
}
