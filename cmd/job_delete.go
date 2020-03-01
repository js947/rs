package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	delete := &cobra.Command{
		Use: "delete <jobid>",
		Short: "Rename job",
		Run: func(cmd *cobra.Command, args []string) {
			err := job_delete(cmd, args[0])
			if err != nil {
				panic(err)
			}
		},
	}
	jobCmd.AddCommand(delete)
}

func job_delete(cmd *cobra.Command, jobid string) error {
	return fmt.Errorf("job_delete not implemented")
}