package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ls = &cobra.Command{
	Use:   "ls",
	Short: "list all the active git profiles",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listing...")
	},
}

func init() {
	rootCmd.AddCommand(ls)
}
