package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "job",
		Short: "job [list,rename,delete]",
	}
	rootCmd.AddCommand(cmd)

	list := &cobra.Command{
		Use: "list",
		Short: "List jobs",
		Run: func(cmd *cobra.Command, args []string) {
			err := jobs(cmd)
			if err != nil {
				panic(err)
			}
		},
	}
	list.Flags().BoolP("all", "a", false, "All jobs (not just owned by me)")
	cmd.AddCommand(list)

	rename := &cobra.Command{
		Use: "rename <jobid> <new name>",
		Short: "Rename job",
		Run: func(cmd *cobra.Command, args []string) {
			err := job_rename(cmd, args[0], args[1])
			if err != nil {
				panic(err)
			}
		},
	}
	cmd.AddCommand(rename)

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
	cmd.AddCommand(delete)
}

func job_rename(cmd *cobra.Command, jobid string, newname string) error {
	return fmt.Errorf("job_rename not implemented")
}

func job_delete(cmd *cobra.Command, jobid string) error {
	return fmt.Errorf("job_delete not implemented")
}