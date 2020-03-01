package cmd

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/js947/rs/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var job = viper.New()

func init() {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Submit job",
		Run:   submit,
	}

	cmd.PersistentFlags().String("config", "rescale", "job config file")
	cmd.PersistentFlags().StringP("path", "p", ".", "path to job")

	cmd.Flags().StringP("core", "c", "", "core type")
	job.BindPFlag("core", cmd.Flags().Lookup("core"))
	cmd.Flags().IntP("numcores", "n", 0, "number of cores")
	job.BindPFlag("numcores", cmd.Flags().Lookup("numcores"))

	rootCmd.AddCommand(cmd)
}

func submit(cmd *cobra.Command, args []string) {
	name, err := cmd.Flags().GetString("config")
	if err != nil {
		panic(err)
	}
	job.SetConfigName(name)

	path, err := cmd.Flags().GetString("path")
	if err != nil {
		panic(err)
	}
	job.AddConfigPath(path)

	job.SetDefault("core", "hpc-3")
	job.SetDefault("numcores", 1)

	err = job.ReadInConfig()
	if err != nil {
		panic(err)
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

	buf := new(bytes.Buffer)
	z := zip.NewWriter(buf)

	if err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			rp, err := filepath.Rel(path, p)
			if err != nil {
				return err
			}
			f, err := z.Create(rp)
			if err != nil {
				return err
			}
			dat, err := ioutil.ReadFile(p)
			if err != nil {
				return err
			}
			nb, err := f.Write([]byte(dat))
			if err != nil {
				return err
			}
			log.Printf("collected input file: %q (%d bytes)\n", rp, nb)
		}
		return nil
	}); err != nil {
		panic(err)
	}
	if err := z.Close(); err != nil {
		panic(err)
	}

	fileinfo, err := api.UploadFile(j.Name+".zip", buf)
	if err != nil {
		panic(err)
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
		in := make([]InputFile, 1)
		in[0] = InputFile{ID: fileinfo.ID}
		ja[i] = JobAnalysis{UseMPI: true, Command: a.Command, Analysis: at, Hardware: ht, InputFiles: in}
	}
	js := Job{Name: j.Name, Analyses: ja}

	jb, err := json.MarshalIndent(js, "", "  ")
	fmt.Printf("%s\n", jb)

	jbuf, err := api.PostJSON("https://platform.rescale.com/api/v2/jobs/", bytes.NewBuffer(jb))
	if err != nil {
		log.Fatal(err)
	}

	var ji struct {
		ID string `json:"id"`
	}
	json.Unmarshal(jbuf, &ji)
	fmt.Printf("created job %s\n", ji.ID)
}
