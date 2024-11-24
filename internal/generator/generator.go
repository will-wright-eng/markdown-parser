package generator

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "github.com/will-wright-eng/parse/internal/parser"
)

type GeneratorOptions struct {
    Overwrite    bool
    FileMode     os.FileMode
    DirMode      os.FileMode
    PreProcess   func(content string, language string) string
    PostProcess  func(path string) error
    SkipPatterns []string
}

func NewDefaultOptions() GeneratorOptions {
    return GeneratorOptions{
        Overwrite:    false,
        FileMode:     0644,
        DirMode:      0755,
        PreProcess:   nil,
        PostProcess:  nil,
        SkipPatterns: []string{},
    }
}

func GenerateFiles(blocks []parser.CodeBlock, baseDir string, opts GeneratorOptions) error {
    for _, block := range blocks {
        if shouldSkipFile(block.FilePath, opts.SkipPatterns) {
            continue
        }

        // Clean and validate the file path
        cleanPath := filepath.Clean(block.FilePath)
        if cleanPath == "" || strings.Contains(cleanPath, "..") {
            continue
        }

        // Create full path
        fullPath := filepath.Join(baseDir, cleanPath)

        // Check if file exists
        if !opts.Overwrite {
            if _, err := os.Stat(fullPath); err == nil {
                continue
            }
        }

        // Create directory structure
        dir := filepath.Dir(fullPath)
        if err := os.MkdirAll(dir, opts.DirMode); err != nil {
            return fmt.Errorf("error creating directory %s: %w", dir, err)
        }

        // Process content
        content := block.Content
        if opts.PreProcess != nil {
            content = opts.PreProcess(content, block.Language)
        }

        // Write file
        if err := os.WriteFile(fullPath, []byte(content), opts.FileMode); err != nil {
            return fmt.Errorf("error writing file %s: %w", fullPath, err)
        }

        // Post-process
        if opts.PostProcess != nil {
            if err := opts.PostProcess(fullPath); err != nil {
                return fmt.Errorf("error post-processing file %s: %w", fullPath, err)
            }
        }
    }
    return nil
}

func shouldSkipFile(path string, patterns []string) bool {
    for _, pattern := range patterns {
        if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
            return true
        }
    }
    return false
}
