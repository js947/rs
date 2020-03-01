package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/js947/rs/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cmd := &cobra.Command{
		Use:   "jobs",
		Short: "List jobs",
		Run: func(cmd *cobra.Command, args []string) {
			err := jobs(cmd)
            if err != nil {
                panic(err)
            }
		},
	}

	cmd.Flags().BoolP("all", "a", false, "All jobs (not just owned by me)")

	rootCmd.AddCommand(cmd)
}

func jobs(cmd *cobra.Command) error {
	type Job struct {
		Name string `json:"name"`
		ID string `json:"id"`
		Owner string `json:"owner"`
	}
	var jobs []Job

	addr := "https://platform.rescale.com/api/v2/jobs/?page_size=20"
	for addr != "" {
		var capi struct {
			Count int
			Next string
			Results []Job
		}
		data, err := api.Get(addr)
		if err != nil {
			return err
		}
		json.Unmarshal(data, &capi)

		addr = capi.Next
		jobs = append(jobs, capi.Results...)

	}

	all_jobs, err := cmd.Flags().GetBool("all")
	if err != nil {
		return err
	}

	if all_jobs {
		f := "%6s\t%24s\t%s\n"
		fmt.Printf(f, "id", "owner", "name")
		for _, j := range jobs {
			fmt.Printf(f, j.ID, j.Owner, j.Name)
		}
	} else {
		f := "%6s\t%s\n"
		fmt.Printf(f, "id", "name")
		for _, j := range jobs {
			if j.Owner == viper.GetString("username") {
				fmt.Printf(f, j.ID, j.Name)
			}
		}
	}

	return nil
}