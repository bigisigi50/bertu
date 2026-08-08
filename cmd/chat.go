package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Codebase Chatter",
	Long:  `Provides an interface to ask questions about your codebase (Mock).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Codebase Chatter: [MOCK] Your codebase consists of a CLI application built with Go and Cobra.")
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
