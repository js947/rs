package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/js947/rs/api"
	wordwrap "github.com/mitchellh/go-wordwrap"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func init() {
	analyses := &cobra.Command{
		Use:   "analyses",
		Short: "Search analysis types (shorthand for 'analysis search')",
		Run:   analysis_search,
	}
	rootCmd.AddCommand(analyses)

	cmd := &cobra.Command{
		Use:   "search",
		Short: "Search analysis types",
		Run:   analysis_search,
	}
	analysisCmd.AddCommand(cmd)
}

type Analysis struct {
	Code        string
	Name        string
	Industries  string
	Description string
	Vendor      string `json:"vendorName"`
}

func analysis_search(cmd *cobra.Command, args []string) {
	addr := "https://platform.rescale.com/api/v2/analyses/?page_size=300"
	var analyses []Analysis

	for addr != "" {
		var capi struct {
			Count   int
			Next    string
			Results []Analysis
		}
		data, err := api.Get(addr)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(data, &capi)

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
		for _, a := range searchResults.Hits {
			var (
				name        string
				description string
			)
			for _, b := range analyses {
				if a.ID == b.Code {
					name = b.Name
					description = b.Description
				}
			}
			fmt.Printf("%20s %s\n", "'" + a.ID + "'", name)
			fmt.Printf("%s\n\n", wordwrap.WrapString(description, 80))
		}
	} else {
		for _, a := range analyses {
			fmt.Printf("%20s%s\n", "'" + a.Code + "'", a.Name)
			fmt.Printf("%s\n\n", wordwrap.WrapString(a.Description, 80))
		}
	}
}
