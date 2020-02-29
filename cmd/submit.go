package cmd

import (
	"log"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var job = viper.New();

func init() {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Submit job",
		Run: submit,
	}

	cmd.PersistentFlags().String("config", "rescale", "job config file")
	cmd.PersistentFlags().StringP("path", "p", ".", "path to job")

	cmd.Flags().StringP("core", "c", "", "core type")
	job.BindPFlag("core", cmd.Flags().Lookup("core"))
	cmd.Flags().IntP("numcores", "n", 0, "number of cores")
	job.BindPFlag("numcores", cmd.Flags().Lookup("numcores"))

	rootCmd.AddCommand(cmd)
}

type Job struct {
	Core string
	NumCores int
	Analysis []AnalysisStep
}
type AnalysisStep struct {
	Software string
	Version string
	Command string
}

func submit(cmd *cobra.Command, args []string) {
	name, err := cmd.Flags().GetString("config")
	if err != nil {
		log.Fatal(err)
	}
	job.SetConfigName(name)

	path, err := cmd.Flags().GetString("path")
	if err != nil {
		log.Fatal(err)
	}
	job.AddConfigPath(path)

	job.SetDefault("core", "onyx")
	job.SetDefault("numcores", 2)

	err = job.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	var j Job
	job.Unmarshal(&j)

	fmt.Printf("core      %s\n", j.Core)
	fmt.Printf("num cores %d\n", j.NumCores)

	for i, a := range j.Analysis {
		fmt.Printf("analysis step %d: software %s\n", i, a.Software)
		fmt.Printf("analysis step %d: version  %s\n", i, a.Version)
		fmt.Printf("analysis step %d: command  %s\n", i, a.Command)
	}
}