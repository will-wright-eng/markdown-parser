package cmd

import (
    "os"

    "github.com/spf13/cobra"
    "parse/internal/logger"
)

var rootCmd = &cobra.Command{
    Use:   "mycli",
    Short: "A CLI tool to generate files and directories from markdown",
    Long: `This CLI tool parses markdown files and generates the corresponding
file and directory structure based on the content.`,
}

func Execute() {
    logger := logger.New()
    if err := rootCmd.Execute(); err != nil {
        logger.Error("Failed to execute command", "error", err)
        os.Exit(1)
    }
}
