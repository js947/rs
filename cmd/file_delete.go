package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rename := &cobra.Command{
		Use:   "delete <jobid> <new name>",
		Short: "Delete file",
		Run: func(cmd *cobra.Command, args []string) {
			err := file_delete(cmd, args[0], args[1])
			if err != nil {
				panic(err)
			}
		},
	}
	jobCmd.AddCommand(rename)
}

func file_delete(cmd *cobra.Command, jobid string, newname string) error {
	return fmt.Errorf("file_rename not implemented")
}
