package parser

import (
    "bufio"
    "fmt"
    "io"
    "path/filepath"
    "regexp"
    "strings"
)

type CodeBlock struct {
    FilePath string
    Content  string
    Language string // Added to support language-specific handling
}

// Parser configuration options
type ParserOptions struct {
    HeaderPrefixes []string        // e.g., ["##", "#", "file:"]
    PathValidators []PathValidator // Custom path validation functions
    StripComments bool            // Whether to strip comments from code blocks
}

type PathValidator func(string) bool

// NewDefaultOptions returns default parser options
func NewDefaultOptions() ParserOptions {
    return ParserOptions{
        HeaderPrefixes: []string{"##", "#", "file:", "path:"},
        PathValidators: []PathValidator{
            func(path string) bool {
                return strings.Contains(path, "/") || strings.Contains(path, ".")
            },
        },
        StripComments: false,
    }
}

func ParseMarkdown(reader io.Reader, opts ParserOptions) ([]CodeBlock, error) {
    var blocks []CodeBlock
    scanner := bufio.NewScanner(reader)

    var currentPath string
    var currentContent strings.Builder
    var currentLanguage string
    inCodeBlock := false

    // Combine all prefixes into a single regex
    prefixPattern := fmt.Sprintf("^(%s)\\s*", strings.Join(opts.HeaderPrefixes, "|"))
    prefixRegex := regexp.MustCompile(prefixPattern)

    for scanner.Scan() {
        line := scanner.Text()
        trimmedLine := strings.TrimSpace(line)

        // Handle file paths
        if !inCodeBlock && trimmedLine != "" && !strings.HasPrefix(trimmedLine, "```") {
            if path := extractFilePath(trimmedLine, prefixRegex); path != "" {
                if isValidPath(path, opts.PathValidators) {
                    if shouldSaveBlock(currentPath, currentContent.String()) {
                        blocks = append(blocks, createCodeBlock(currentPath, currentContent.String(), currentLanguage))
                        currentContent.Reset()
                    }
                    currentPath = filepath.Clean(path)
                }
            }
            continue
        }

        // Handle code block markers
        if strings.HasPrefix(trimmedLine, "```") {
            if inCodeBlock {
                // End of code block
                if shouldSaveBlock(currentPath, currentContent.String()) {
                    blocks = append(blocks, createCodeBlock(currentPath, currentContent.String(), currentLanguage))
                    currentContent.Reset()
                }
                inCodeBlock = false
                currentLanguage = ""
            } else {
                // Start of code block
                inCodeBlock = true
                currentLanguage = extractLanguage(trimmedLine)
            }
            continue
        }

        // Add content if we're in a code block
        if inCodeBlock {
            if opts.StripComments {
                line = stripComments(line, currentLanguage)
            }
            if line != "" {
                currentContent.WriteString(line + "\n")
            }
        }
    }

    // Handle the last block
    if shouldSaveBlock(currentPath, currentContent.String()) {
        blocks = append(blocks, createCodeBlock(currentPath, currentContent.String(), currentLanguage))
    }

    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("error scanning markdown: %w", err)
    }

    return blocks, nil
}

// Helper functions
func extractFilePath(line string, prefixRegex *regexp.Regexp) string {
    // Remove any prefix and clean the path
    return strings.TrimSpace(prefixRegex.ReplaceAllString(line, ""))
}

func isValidPath(path string, validators []PathValidator) bool {
    if len(validators) == 0 {
        return true
    }
    for _, validator := range validators {
        if validator(path) {
            return true
        }
    }
    return false
}

func shouldSaveBlock(path, content string) bool {
    return path != "" && strings.TrimSpace(content) != ""
}

func createCodeBlock(path, content, language string) CodeBlock {
    return CodeBlock{
        FilePath: path,
        Content:  strings.TrimSpace(content) + "\n",
        Language: language,
    }
}

func extractLanguage(line string) string {
    parts := strings.Fields(strings.TrimPrefix(line, "```"))
    if len(parts) > 0 {
        return parts[0]
    }
    return ""
}

func stripComments(line, language string) string {
    line = strings.TrimSpace(line)
    switch language {
    case "go":
        if strings.HasPrefix(line, "//") {
            return ""
        }
    case "python":
        if strings.HasPrefix(line, "#") {
            return ""
        }
    case "javascript", "typescript":
        if strings.HasPrefix(line, "//") {
            return ""
        }
    }
    return line
}
