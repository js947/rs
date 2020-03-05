package cmd

import (
	"fmt"
	"github.com/js947/rs/api"
	"github.com/spf13/cobra"
)

func init() {
	delete := &cobra.Command{
		Use:   "delete <fileid> [<fileid>...]",
		Short: "Delete file(s)",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := file_delete(args)
			if err != nil {
				panic(err)
			}
		},
	}
	fileCmd.AddCommand(delete)
}

func file_delete(fileids []string) error {
	for _, fileid := range fileids {
		addr := fmt.Sprintf("https://platform.rescale.com/api/v2/files/%s/", fileid)
		_, err :=  api.Delete(addr)
		if err != nil {
			return err
		}
	}
	return nil
}
