package cmd

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/bigisigi50/bertu/internal/ai"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Smart Git Assistant",
	Long:  `Runs git diff and automatically generates a conventional commit message using Gemini.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Get git diff
		diffCmd := exec.Command("git", "diff", "--cached")
		output, err := diffCmd.Output()
		if err != nil {
			fmt.Printf("Error running git diff: %v\n", err)
			return
		}

		diffString := string(output)
		if strings.TrimSpace(diffString) == "" {
			fmt.Println("No staged changes found. Run 'git add' first.")
			return
		}

		// 2. Prepare the prompt
		prompt := fmt.Sprintf(`You are an expert developer. Look at this git diff and write a single conventional commit message summarizing the changes. Do not include any other text, just the commit message itself.

Diff:
%s`, diffString)

		fmt.Println("Generating commit message...")

		// 3. Call Gemini
		ctx := context.Background()
		response, err := ai.GenerateContent(ctx, prompt)
		if err != nil {
			fmt.Printf("Error generating commit message: %v\n", err)
			return
		}

		fmt.Printf("\nGenerated commit message:\n\n%s\n", strings.TrimSpace(response))
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)
}
