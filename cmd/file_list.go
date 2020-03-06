package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/js947/rs/api"
	"github.com/spf13/cobra"
)

func init() {
	files := &cobra.Command{
		Use:   "files",
		Short: "List files (shorthand for 'file list')",
		Run: func(cmd *cobra.Command, args []string) {
			err := file_list(cmd)
			if err != nil {
				panic(err)
			}
		},
	}
	rootCmd.AddCommand(files)

	list := &cobra.Command{
		Use:   "list",
		Short: "List files",
		Run: func(cmd *cobra.Command, args []string) {
			err := file_list(cmd)
			if err != nil {
				panic(err)
			}
		},
	}
	fileCmd.AddCommand(list)
}

type File struct {
	TypeID       int       `json:"typeId"`
	Name         string    `json:"name"`
	Date         time.Time `json:"dateUploaded"`
	RelativePath string    `json:"relativePath"`
	URL          string    `json:"downloadUrl"`
	Size         int       `json:"decryptedSize"`
	Path         string    `json:"path"`
	ID           string    `json:"id"`
	MD5          string    `json:"md5"`
	IsDeleted    bool      `json:"isDelelted"`
}

func file_list(cmd *cobra.Command) error {
	var files []File
	var count int
	addr := "https://platform.rescale.com/api/v3/files/?page=1&owner=1&page_size=20&ordering=date_inserted"
	for addr != "" {
		fmt.Printf("get %s\n", addr)
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

func (f *File) TypeStr() string {
	if f.TypeID == 1 {
		return "input"
	}
	if f.TypeID == 5 {
		return "output"
	}
	return fmt.Sprintf("%d", f.TypeID)
}
