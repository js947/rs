package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/js947/rs/api"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "cores [application version]",
		Short: "List core types",
		Run:   func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				allcores()
			} else {
				cores_for_application(args[0], args[1])
			}
		},
	}
	rootCmd.AddCommand(cmd)
}

type Core struct {
	Code      string
	Name      string
	Processor string `json:"processorInfo"`
}
func get_coretypes() []Core {
	var cores []Core

	addr := "https://platform.rescale.com/api/v2/coretypes/?page_size=20"
	for addr != "" {
		var capi struct {
			Count   int
			Next    string
			Results []Core
		}
		json.Unmarshal(api.Get(addr), &capi)

		addr = capi.Next
		cores = append(cores, capi.Results...)
	}
	return cores
}

func allcores() {
	f := "%12s\t%12s\t%s\n"
	fmt.Printf(f, "code", "name", "description")
	for _, v := range get_coretypes() {
		fmt.Printf(f, v.Code, v.Name, v.Processor)
	}
}

func cores_for_application(app string, version string) {
	type ApplicationVersion struct {
		Id string
		AllowedCoreTypes []string `json:"allowedCoreTypes"`
		Version string
		Code string `json:"versionCode"`
	}
	var ad struct {
		Code string
		Description string
		Versions []ApplicationVersion
	}

	addr := fmt.Sprintf("https://platform.rescale.com/api/v2/analyses/%s/", app)
	json.Unmarshal(api.Get(addr), &ad)

	for _, v := range ad.Versions {
		if v.Code == version {
			for _, c := range v.AllowedCoreTypes {
				fmt.Printf("%12s\n", c)
			}
			break
		}
	}
}