package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rs",
	Short: "rs is a simple cli for the rescale scaleX platform",
	Long: `rs is a simple cli for the rescale scaleX platform made by js947. 
See http://github.com/js947/rs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hi")
	},
  }
  
  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }