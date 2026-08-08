package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/bigisigi50/bertu/internal/ai"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Terminal Error Explainer",
	Long:  `Reads standard input and uses Gemini to explain the error and provide a fix.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Read from stdin
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			fmt.Println("No piped input detected. Usage: <failing_command> | bertu explain")
			return
		}

		inputBytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Printf("Error reading stdin: %v\n", err)
			return
		}

		errorText := string(inputBytes)
		if strings.TrimSpace(errorText) == "" {
			fmt.Println("Piped input is empty.")
			return
		}

		// 2. Prepare the prompt
		prompt := fmt.Sprintf(`The following is an error output from a terminal command. Explain why it failed and provide the exact steps or commands to fix it in a concise manner:

%s`, errorText)

		fmt.Println("Analyzing error...")

		// 3. Call Gemini
		ctx := context.Background()
		response, err := ai.GenerateContent(ctx, prompt)
		if err != nil {
			fmt.Printf("Error explaining error: %v\n", err)
			return
		}

		fmt.Printf("\nExplanation & Fix:\n\n%s\n", strings.TrimSpace(response))
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)
}
