package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/js947/rs/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	jobs := &cobra.Command{
		Use:   "jobs",
		Short: "List jobs (shorthand for 'job list')",
		Run: func(cmd *cobra.Command, args []string) {
			err := job_list(cmd)
			if err != nil {
				panic(err)
			}
		},
	}
	rootCmd.AddCommand(jobs)

	list := &cobra.Command{
		Use:   "list",
		Short: "List jobs",
		Run: func(cmd *cobra.Command, args []string) {
			err := job_list(cmd)
			if err != nil {
				panic(err)
			}
		},
	}
	jobCmd.AddCommand(list)

	joblist_flags := func(cmd *cobra.Command) {
		cmd.Flags().BoolP("all", "a", false, "All jobs (not just owned by me)")
	}
	joblist_flags(jobs)
	joblist_flags(list)
}

func job_list(cmd *cobra.Command) error {
	type Job struct {
		Name      string `json:"name"`
		ID        string `json:"id"`
		Owner     string `json:"owner"`
		JobStatus struct {
			State string `json:"content"`
		} `json:"jobStatus"`
		ClusterStatus struct {
			State string `json:"content"`
		} `json:"clusterStatusDisplay"`
		Date time.Time `json:"dateInserted"`
	}
	var jobs []Job

	addr := "https://platform.rescale.com/api/v2/jobs/?page_size=20"
	for addr != "" {
		var capi struct {
			Count   int
			Next    string
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
		f := "%6s\t%24s\t%24s\t%s\n"
		fmt.Printf(f, "id", "owner", "status", "name")
		for _, j := range jobs {
			fmt.Printf(f, j.ID, j.Owner, j.JobStatus.State+"/"+j.ClusterStatus.State, j.Name)
		}
	} else {
		f := "%6s\t%24s\t%16s\t%s\n"
		fmt.Printf(f, "id", "date", "status", "name")
		for _, j := range jobs {
			if j.Owner == viper.GetString("username") {
				fmt.Printf(f,
					j.ID,
					j.Date.Format("Mon Jan _2 15:04 2006"),
					j.JobStatus.State+"/"+j.ClusterStatus.State,
					j.Name)
			}
		}
	}

	return nil
}
