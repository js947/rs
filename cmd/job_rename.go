package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rename := &cobra.Command{
		Use:   "rename <jobid> <new name>",
		Short: "Rename job",
		Run: func(cmd *cobra.Command, args []string) {
			err := job_rename(cmd, args[0], args[1])
			if err != nil {
				panic(err)
			}
		},
	}
	jobCmd.AddCommand(rename)
}

func job_rename(cmd *cobra.Command, jobid string, newname string) error {
	return fmt.Errorf("job_rename not implemented")
}
