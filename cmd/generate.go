// cmd/generate.go
package cmd

import (
    "fmt"
    "os"

    "parse/internal/generator"
    "parse/internal/parser"

    "github.com/spf13/cobra"
)

var (
    inputFile     string
    outputDir     string
    overwrite     bool
    stripComments bool
    skipPatterns  []string
)

var generateCmd = &cobra.Command{
    Use:   "generate",
    Short: "Generate project files from markdown",
    Long: `Generate project files and directory structure from a markdown file.
The markdown file can contain file paths in various formats:
- As headers (## path/to/file.txt)
- With prefixes (file: path/to/file.txt)
- As bare paths (path/to/file.txt)

Each file path should be followed by a code block containing the file content.`,
    RunE: func(cmd *cobra.Command, args []string) error {
        // Open input file
        file, err := os.Open(inputFile)
        if err != nil {
            return fmt.Errorf("error opening input file: %w", err)
        }
        defer file.Close()

        // Configure parser options
        parserOpts := parser.NewDefaultOptions()
        parserOpts.StripComments = stripComments

        // Parse markdown
        blocks, err := parser.ParseMarkdown(file, parserOpts)
        if err != nil {
            return fmt.Errorf("error parsing markdown: %w", err)
        }

        // Configure generator options
        genOpts := generator.NewDefaultOptions()
        genOpts.Overwrite = overwrite
        genOpts.SkipPatterns = skipPatterns

        // Generate files
        if err := generator.GenerateFiles(blocks, outputDir, genOpts); err != nil {
            return fmt.Errorf("error generating files: %w", err)
        }

        fmt.Printf("Successfully generated files in %s\n", outputDir)
        return nil
    },
}

func init() {
    generateCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input markdown file (required)")
    generateCmd.Flags().StringVarP(&outputDir, "output", "o", "tmp", "Output directory (default: ./tmp)")
    generateCmd.Flags().BoolVarP(&overwrite, "force", "f", false, "Overwrite existing files")
    generateCmd.Flags().BoolVar(&stripComments, "strip-comments", false, "Strip comments from code blocks")
    generateCmd.Flags().StringSliceVar(&skipPatterns, "skip", []string{}, "Patterns of files to skip (e.g., *.tmp)")
    generateCmd.MarkFlagRequired("input")
    rootCmd.AddCommand(generateCmd)
}
