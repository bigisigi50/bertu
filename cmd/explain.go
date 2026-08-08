package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Terminal Error Explainer",
	Long:  `Reads standard input and provides an explanation of the error (Mock).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Terminal Error Explainer: [MOCK] It seems you are missing a dependency. Run 'npm install' to fix this.")
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)
}
