package cmd

import (
	"fmt"

	"github.com/gobuffalo/packr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init AwesomeProjectName",
	Short: "Initialize your new cosmos zone",
	Run: func(cmd *cobra.Command, args []string) {
		box := packr.NewBox("../template")
		gopkg := box.Walk(func(path string, file packr.File) error {
			fmt.Println(path)
			return nil
		})
		fmt.Println(gopkg)
		fmt.Println("Initialized a new project, happy hacking!")
	},
}
