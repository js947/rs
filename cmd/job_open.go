package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
    "github.com/pkg/browser"
)

func init() {
	delete := &cobra.Command{
		Use:   "open <jobid>",
		Short: "Open job in browser",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := job_open(cmd, args[0])
			if err != nil {
				panic(err)
			}
		},
	}
	jobCmd.AddCommand(delete)
}

func job_open(cmd *cobra.Command, jobid string) error {
	platform := "https://platform.rescale.com"
	return browser.OpenURL(fmt.Sprintf("%s/jobs/%s/", platform, jobid))
}
