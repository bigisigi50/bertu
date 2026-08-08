package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Smart Git Assistant",
	Long:  `Runs git diff and automatically generates a conventional commit message (Mock).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Smart Git Assistant: [MOCK] Generated commit message: 'feat: add new CLI scaffolding'")
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)
}
