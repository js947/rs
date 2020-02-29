package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
)

type Job struct {
	Core     string
	NumCores int
	Analysis []AnalysisStep
}
type AnalysisStep struct {
	Software string
	Version  string
	Command  string
}

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "submit job",
	Run: func(cmd *cobra.Command, args []string) {
		doSubmit()
	},
}

func init() {
	rootCmd.AddCommand(submitCmd)
}

func doSubmit() {
	job_file, err := os.Open("rescale.yaml")
	if err != nil {
		log.Fatal(err)
	}

	job_raw, err := ioutil.ReadAll(job_file)
	if err != nil {
		log.Fatal(err)
	}

	var job Job
	err = yaml.Unmarshal([]byte(job_raw), &job)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(job)
}