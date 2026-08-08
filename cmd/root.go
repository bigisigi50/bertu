package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bertu",
	Short: "Bertu is a CLI tool for AI tooling",
	Long:  `Bertu provides AI-assisted development tools (Smart Git Assistant, Error Explainer, Codebase Chatter).`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
