package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
  "github.com/gobuffalo/packr"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "init AwesomeProjectName [flags]",
  Short: "Initialize your awesome cosmos-sdk zone", 
  Run: func(cmd *cobra.Command, args []string) {
    box := packr.NewBox("../template")
    gopkg := box.String("Gopkg.toml")
    fmt.Println(gopkg)
    fmt.Println("Initialized a new project, happy hacking!")
  },
}
