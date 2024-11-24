package main

import (
	"github.com/spf13/cobra"
	"strings"
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
			tmpPrefix, _ := cmd.Flags().GetString("tmpPrefix")
			if !strings.HasSuffix(tmpPrefix, "/") {
				tmpPrefix += "/"
			}
			err := parseMarkdownWithPrefix(file, tmpPrefix)
			if err != nil {
				log.Fatalf("Error parsing markdown: %v", err)
			}
		},
	}

	// Add a flag for tmpPrefix with a default value of "tmp/"
	parseCmd.Flags().StringP("tmpPrefix", "t", "tmp/", "Temporary file prefix")

	rootCmd.AddCommand(parseCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing root command: %v", err)
		os.Exit(1)
	}
}
