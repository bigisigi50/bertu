package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bigisigi50/bertu/internal/ai"
	"github.com/philippgille/chromem-go"
	"github.com/spf13/cobra"
)

var initFlag bool

var chatCmd = &cobra.Command{
	Use:   "chat [question]",
	Short: "Codebase Chatter",
	Long:  `Ask questions about your codebase. Use --init to index your local Go files first.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		dbPath := ".chromem"

		// Ensure API key is set for embeddings and chat
		if os.Getenv("GEMINI_API_KEY") == "" {
			fmt.Println("GEMINI_API_KEY environment variable is not set")
			return
		}

		if initFlag {
			fmt.Println("Initializing codebase index...")
			err := indexCodebase(ctx, dbPath)
			if err != nil {
				fmt.Printf("Error indexing codebase: %v\n", err)
			}
			return
		}

		if len(args) == 0 {
			fmt.Println("Please provide a question. Example: bertu chat 'How does the chat command work?'")
			return
		}
		question := strings.Join(args, " ")

		fmt.Println("Querying codebase...")
		err := answerQuestion(ctx, dbPath, question)
		if err != nil {
			fmt.Printf("Error answering question: %v\n", err)
		}
	},
}

func indexCodebase(ctx context.Context, dbPath string) error {
	// Initialize Chromem DB
	db, err := chromem.NewPersistentDB(dbPath, false)
	if err != nil {
		return err
	}

	// Create collection (using default local embedding function for simplicity if supported, or custom)
	// Note: chromem-go requires an embedding function. We'll use a dummy/basic one or OpenAI if default not available.
	// We will just use the default OpenAI embedding function with a dummy key for local testing or custom.
	// Actually, let's just use the default embedding function provided by chromem.
	collection, err := db.GetOrCreateCollection("codebase", nil, nil)
	if err != nil {
		return err
	}

	// Walk directory and index .go files
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") && !strings.Contains(path, "vendor") {
			content, err := os.ReadFile(path)
			if err != nil {
				return nil
			}

			// Add document
			doc := chromem.Document{
				ID:      path,
				Content: string(content),
				Metadata: map[string]string{
					"file": path,
				},
			}
			err = collection.AddDocument(ctx, doc)
			if err != nil {
				fmt.Printf("Failed to index %s: %v\n", path, err)
			} else {
				fmt.Printf("Indexed %s\n", path)
			}
		}
		return nil
	})

	if err == nil {
		fmt.Println("Index initialization complete!")
	}
	return err
}

func answerQuestion(ctx context.Context, dbPath string, question string) error {
	db, err := chromem.NewPersistentDB(dbPath, false)
	if err != nil {
		return fmt.Errorf("failed to open db (did you run with --init first?): %w", err)
	}

	collection := db.GetCollection("codebase", nil)
	if collection == nil {
		return fmt.Errorf("collection not found. Please run with --init first")
	}

	// Search for relevant files
	res, err := collection.Query(ctx, question, 5, nil, nil)
	if err != nil {
		return err
	}

	var contextStrings []string
	for _, doc := range res {
		contextStrings = append(contextStrings, fmt.Sprintf("File: %s\nContent:\n%s", doc.ID, doc.Content))
	}

	if len(contextStrings) == 0 {
		fmt.Println("No relevant code found in the index.")
		return nil
	}

	// Prepare prompt with context
	prompt := fmt.Sprintf(`Answer the user's question based on the following code snippets from their codebase.

Codebase Context:
%s

User Question: %s`, strings.Join(contextStrings, "\n\n---\n\n"), question)

	// Call Gemini
	response, err := ai.GenerateContent(ctx, prompt)
	if err != nil {
		return err
	}

	fmt.Printf("\nAnswer:\n\n%s\n", strings.TrimSpace(response))
	return nil
}

func init() {
	chatCmd.Flags().BoolVar(&initFlag, "init", false, "Initialize and index the local codebase")
	rootCmd.AddCommand(chatCmd)
}
