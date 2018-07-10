package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "init AwesomeProjectName [flags]",
  Short: "Initialize your awesome cosmos-sdk zone", 
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Initialized a new project, happy hacking!")
  },
}
