package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fooCmd = &cobra.Command{
	Use:   "foo",
	Short: "Prints hello world",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world again")
	},
}

func init() {
	rootCmd.AddCommand(fooCmd)
}
