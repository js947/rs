package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/js947/rs/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var job = viper.New()

func init() {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Submit job",
		Run: func(cmd *cobra.Command, args []string) {
			err := submit(cmd)
			if err != nil {
				panic(err)
			}
		},
	}

	cmd.PersistentFlags().String("config", "rescale", "job config file")
	cmd.PersistentFlags().StringP("path", "p", ".", "path to job")

	cmd.Flags().StringP("core", "c", "", "core type")
	job.BindPFlag("core", cmd.Flags().Lookup("core"))
	cmd.Flags().IntP("numcores", "n", 0, "number of cores")
	job.BindPFlag("numcores", cmd.Flags().Lookup("numcores"))
	cmd.Flags().String("name", "", "job name")
	job.BindPFlag("name", cmd.Flags().Lookup("name"))

	cmd.Flags().BoolP("watch", "w", false, "watch log file")
	cmd.Flags().BoolP("sync", "s", false, "sync output files")

	rootCmd.AddCommand(cmd)
}

func submit(cmd *cobra.Command) error {
	name, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}
	job.SetConfigName(name)

	path, err := cmd.Flags().GetString("path")
	if err != nil {
		return err
	}
	job.AddConfigPath(path)

	job.SetDefault("name", "Untitled Job")
	job.SetDefault("core", "hpc-3")
	job.SetDefault("numcores", 1)

	err = job.ReadInConfig()
	if err != nil {
		return err
	}

	var j struct {
		Name     string
		Core     string
		NumCores int
		Analysis []struct {
			Software string
			Version  string
			Command  string
		}
	}
	job.Unmarshal(&j)

	fmt.Printf("name      %s\n", j.Name)
	fmt.Printf("core      %s\n", j.Core)
	fmt.Printf("num cores %d\n", j.NumCores)

	for i, a := range j.Analysis {
		fmt.Printf("analysis step %d: software %s\n", i, a.Software)
		fmt.Printf("analysis step %d: version  %s\n", i, a.Version)
		fmt.Printf("analysis step %d: command  %s\n", i, a.Command)
	}

	files, err := do_upload_dir(path)
	if err != nil {
		return err
	}

	type AnalysisType struct {
		Code    string `json:"code"`
		Version string `json:"version"`
	}
	type HardwareType struct {
		CoresPerSlot int    `json:"coresPerSlot"`
		Slots        int    `json:"slots"`
		CoreType     string `json:"coreType"`
	}
	type InputFile struct {
		ID string `json:"id"`
	}
	type JobAnalysis struct {
		UseMPI     bool         `json:"useMpi"`
		Command    string       `json:"command"`
		Analysis   AnalysisType `json:"analysis"`
		Hardware   HardwareType `json:"hardware"`
		InputFiles []InputFile  `json:"inputFiles"`
	}
	type Job struct {
		Name     string        `json:"name"`
		Analyses []JobAnalysis `json:"jobanalyses"`
	}
	ja := make([]JobAnalysis, len(j.Analysis))
	for i, a := range j.Analysis {
		at := AnalysisType{Code: a.Software, Version: a.Version}
		ht := HardwareType{CoresPerSlot: j.NumCores, Slots: 1, CoreType: j.Core}
		in := make([]InputFile, len(files))
		for j, f := range files {
			in[j] = InputFile{ID: f.ID}
		}
		ja[i] = JobAnalysis{UseMPI: true, Command: a.Command, Analysis: at, Hardware: ht, InputFiles: in}
	}
	js := Job{Name: j.Name, Analyses: ja}

	jb, err := json.MarshalIndent(js, "", "  ")
	if err != nil {
		return err
	}
	jbuf, err := api.Post("v2/jobs/", bytes.NewBuffer(jb))
	if err != nil {
		return err
	}

	var ji struct {
		ID string `json:"id"`
	}
	json.Unmarshal(jbuf, &ji)
	fmt.Printf("create job %s - %s\n", ji.ID, j.Name)

	_, err = api.Post(fmt.Sprintf("v2/jobs/%s/submit/", ji.ID), bytes.NewBuffer([]byte("")))
	if err != nil {
		return err
	}
	fmt.Printf("submit job %s - %s\n", ji.ID, j.Name)

	watch, err := cmd.Flags().GetBool("watch")
	if err != nil {
		return err
	}
	if watch {
		return fmt.Errorf("watching job not implemented")
	}

	sync, err := cmd.Flags().GetBool("sync")
	if err != nil {
		return err
	}
	if sync {
		return fmt.Errorf("syncing job output not implemented")
	}
	return nil
}
