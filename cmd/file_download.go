package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/js947/rs/api"
	"github.com/spf13/cobra"
)

func init() {
	download := &cobra.Command{
		Use:   "download",
		Short: "Download file",
		Run: func(cmd *cobra.Command, args []string) {
			err := file_download(args)
			if err != nil {
				panic(err)
			}
		},
	}
	fileCmd.AddCommand(download)
}

func file_download(fileids []string) error {
	for _, fileid := range fileids {
		addr := fmt.Sprintf("https://platform.rescale.com/api/v2/files/%s/", fileid)
		data, err := api.Get(addr)
		if err != nil {
			return err
		}
		var file File
		json.Unmarshal(data, &file)

		name := file.Name

		addr = fmt.Sprintf("https://platform.rescale.com/api/v2/files/%s/contents/", fileid)
		data, err = api.Get(addr)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(name, data, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
