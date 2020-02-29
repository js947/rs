package cmd

import (
	"os"
	"fmt"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/js947/rs/api"
)

func init() {
	cmd := &cobra.Command{
		Use:   "analyses",
		Short: "List analisys types",
		Run: analyses,
	}
	rootCmd.AddCommand(cmd)
}

type Analysis struct {
	Code string
	Name string
	Description string
}

func analyses(cmd *cobra.Command, args []string) {
	addr := "https://platform.rescale.com/api/v2/analyses/?page_size=100"
	var analyses []Analysis

	for addr != "" {
		var capi struct {
			Count int
			Next string
			Results []Analysis
		}
		json.Unmarshal(api.Get(addr), &capi)

		fmt.Fprintf(os.Stderr, "%d, %s\n", capi.Count, capi.Next)

		addr = capi.Next
		analyses = append(analyses, capi.Results...)
	}

	f := "%12s\t%12s\t%s\n"
	fmt.Printf(f, "code", "name", "description")
	for _, v := range analyses {
		fmt.Printf(f, v.Code, v.Name, v.Description)
	}
}