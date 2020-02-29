package cmd

import (
	"fmt"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/js947/rs/api"
)

func init() {
	cmd := &cobra.Command{
		Use:   "cores",
		Short: "List core types",
		Run: cores,
	}
	rootCmd.AddCommand(cmd)
}

type Core struct {
	Code string
	Name string
	Processor string `json:"processorInfo"`
}

func cores(cmd *cobra.Command, args []string) {
	addr := "https://platform.rescale.com/api/v2/coretypes/?page_size=20"
	var cores []Core

	for addr != "" {
		var capi struct {
			Count int
			Next string
			Results []Core
		}
		json.Unmarshal(api.Get(addr), &capi)

		addr = capi.Next
		cores = append(cores, capi.Results...)
	}

	f := "%12s\t%12s\t%s\n"
	fmt.Printf(f, "code", "name", "description")
	for _, v := range cores {
		fmt.Printf(f, v.Code, v.Name, v.Processor)
	}
}