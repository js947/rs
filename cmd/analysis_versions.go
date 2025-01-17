package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/js947/rs/api"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "versions application_code",
		Short: "List versions of a given application",
		Run: func(cmd *cobra.Command, args []string) {
			err := versions(args[0])
			if err != nil {
				panic(err)
			}
		},
	}
	analysisCmd.AddCommand(cmd)
}

func versions(app string) error {
	type ApplicationVersion struct {
		Id               string
		AllowedCoreTypes []string `json:"allowedCoreTypes"`
		Version          string
		Code             string `json:"versionCode"`
	}
	var ad struct {
		Code        string
		Description string
		Versions    []ApplicationVersion
	}

	addr := fmt.Sprintf("https://platform.rescale.com/api/v2/analyses/%s/", app)
	body, err := api.Get(addr)
	if err != nil {
		return err
	}
	json.Unmarshal(body, &ad)

	for _, v := range ad.Versions {
		fmt.Printf("%20s\t%s\n", v.Version, v.Code)
	}
	return nil
}
