package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/js947/rs/api"
	"github.com/spf13/cobra"
)

func init() {
	cat := &cobra.Command{
		Use:   "cat",
		Short: "Read file content",
		Run: func(cmd *cobra.Command, args []string) {
			err := file_cat(args[0])
			if err != nil {
				panic(err)
			}
		},
	}
	fileCmd.AddCommand(cat)
}

func file_cat(fileid string) error {
	addr := fmt.Sprintf("https://platform.rescale.com/api/v2/files/%s/lines/", fileid)
	data, err := api.Get(addr)
	if err != nil {
		return err
	}

	var filecontents struct {
		Lines []string `json:"lines"`
	}
	json.Unmarshal(data, &filecontents)

	for _, line := range filecontents.Lines {
		fmt.Printf("%s", line)
	}
	return nil
}
