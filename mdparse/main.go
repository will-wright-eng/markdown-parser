package main

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "markdown-parser",
		Short: "Markdown Parser CLI",
		Run: func(cmd *cobra.Command, args []string) {
			// This will be executed if no subcommand is given
			log.Println("Use 'markdown-parser parse <file>' to parse a markdown file.")
		},
	}

	var parseCmd = &cobra.Command{
		Use:   "parse [file]",
		Short: "Parse a markdown file",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			file := args[0]
			err := parseMarkdown(file)
			if err != nil {
				log.Fatalf("Error parsing markdown: %v", err)
			}
		},
	}

	rootCmd.AddCommand(parseCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing root command: %v", err)
		os.Exit(1)
	}
}

