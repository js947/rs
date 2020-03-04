package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/js947/rs/api"
	"github.com/spf13/cobra"
)

func init() {
	upload := &cobra.Command{
		Use:   "upload <files...>",
		Short: "Uploads files or directories",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := file_upload(args)
			if err != nil {
				panic(err)
			}
		},
	}
	fileCmd.AddCommand(upload)
}

func do_upload(name string, path string) error {
	fmt.Printf("upload %s from %s\n", name, path)
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = api.UploadFile(name, bytes.NewBuffer(dat))
	if err != nil {
		return err
	}
	return nil
}
func do_upload_dir(path string) error {
	return filepath.Walk(path, func (path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		rpath, err := filepath.Rel(path, path)
		if err != nil {
			return err
		}
		return do_upload(rpath, path)
	})
}

func file_upload(paths []string) error {
	for _, path := range paths {
		file, err := os.Stat(path)
		if err != nil {
			return err
		}

		if file.IsDir() {
			err = do_upload_dir(path)
		} else {
			err = do_upload(path, path)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
