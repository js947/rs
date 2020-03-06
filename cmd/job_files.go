package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/js947/rs/api"
	"github.com/spf13/cobra"
)

func init() {
	files := &cobra.Command{
		Use:   "files",
		Short: "List files associated to a job",
		Run: func(cmd *cobra.Command, args []string) {
			err := job_file_list(cmd, args[0])
			if err != nil {
				panic(err)
			}
		},
	}
	jobCmd.AddCommand(files)
}

func job_file_list(cmd *cobra.Command, jobid string) error {
	var files []File
	var count int
	addr := fmt.Sprintf("https://platform.rescale.com/api/v2/jobs/%s/files/?page_size=200", jobid)
	for addr != "" {
		var capi struct {
			Count   int
			Next    string
			Results []File
		}
		data, err := api.Get(addr)
		if err != nil {
			return err
		}
		json.Unmarshal(data, &capi)

		addr = capi.Next
		files = append(files, capi.Results...)
		count = capi.Count
	}

	f := "%12s\t%12s\t%24s\t%s\n"
	fmt.Printf(f, "ID", "type", "name", "name")
	for _, file := range files {
		if file.IsDeleted {
			continue
		}
		fmt.Printf(f, file.ID, file.TypeStr(), file.Name, file.Path)
	}
	fmt.Printf("%d files\n", count)

	return nil
}
