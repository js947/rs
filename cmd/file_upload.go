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

func do_upload(name string, path string) (*api.FileInfo, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	file, err := api.UploadFile(name, bytes.NewBuffer(dat))
	if err != nil {
		return nil, err
	}
	fmt.Printf("upload %s %s from %s\n", file.ID, name, path)
	return file, nil
}
func do_upload_dir(path string) ([]*api.FileInfo, error) {
	fmt.Printf("uploading dir %s\n", path)
	var files []*api.FileInfo
	err := filepath.Walk(path, func (p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		rp, err := filepath.Rel(path, p)
		if err != nil {
			return err
		}
		file, err := do_upload(rp, p)
		if err != nil {
			return err
		}
		files = append(files, file)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func file_upload(paths []string) error {
	for _, path := range paths {
		file, err := os.Stat(path)
		if err != nil {
			return err
		}

		if file.IsDir() {
			_, err = do_upload_dir(path)
		} else {
			_, err = do_upload(path, path)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
