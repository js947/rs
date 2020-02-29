package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/blevesearch/bleve"
	"github.com/js947/rs/api"
	"github.com/spf13/cobra"
	"os"
	"strings"
	wordwrap "github.com/mitchellh/go-wordwrap"
)

func init() {
	cmd := &cobra.Command{
		Use:   "analyses",
		Short: "List analysis types",
		Run:   analyses,
	}
	rootCmd.AddCommand(cmd)
}

type Analysis struct {
	Code        string
	Name        string
	Industries  string
	Description string
	Vendor      string `json:"vendorName"`
}

func analyses(cmd *cobra.Command, args []string) {
	addr := "https://platform.rescale.com/api/v2/analyses/?page_size=300"
	var analyses []Analysis

	for addr != "" {
		var capi struct {
			Count   int
			Next    string
			Results []Analysis
		}
		json.Unmarshal(api.Get(addr), &capi)

		addr = capi.Next
		analyses = append(analyses, capi.Results...)
	}

	if len(args) > 0 {
		mapping := bleve.NewIndexMapping()

		dir, err := ioutil.TempDir(os.TempDir(), "analyses.search")
		if err != nil {
			log.Fatal(err)
		}
		defer os.RemoveAll(dir)

		index, err := bleve.New(dir, mapping)
		if err != nil {
			log.Fatal(err)
		}
		for _, a := range analyses {
			index.Index(a.Code, a.Description)
		}

		count, err := index.DocCount()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d analyses indexed\n", count)

		query := bleve.NewMatchQuery(strings.Join(args, " "))
		search := bleve.NewSearchRequest(query)
		searchResults, err := index.Search(search)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d matches\n", len(searchResults.Hits))
		for i, a := range searchResults.Hits {
			var (
				name string
				description string
			)
			for _, b := range analyses {
				if a.ID == b.Code {
					name = b.Name
					description = b.Description
				}
			}
			fmt.Printf("%3d %s\n", i, a.ID) 
			fmt.Printf("%s\n%s\n\n", name, wordwrap.WrapString(description, 80)) 
		}
	} else {
		for i, a := range analyses {
			fmt.Printf("%3d %s\n", i, a.Code) 
			fmt.Printf("%s\n%s\n\n", a.Name, wordwrap.WrapString(a.Description, 80)) 
		}
	}
}
