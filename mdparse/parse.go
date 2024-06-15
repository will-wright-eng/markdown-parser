package main

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func parseMarkdownWithPrefix(file string, tmpPrefix string) error {
	log.Printf("Starting to parse the markdown file: %s", file)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("Error reading file: %s", err)
		return fmt.Errorf("failed to read file: %w", err)
	}

	md := goldmark.New()
	doc := md.Parser().Parse(text.NewReader(content))

	var filename string
	err = ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if n.Kind() == ast.KindText {
			textNode := n.(*ast.Text)
			log.Printf("Bold text found: %s", textNode.Text(content))
			textContent := string(textNode.Text(content))
			filename = extractFilename(textContent)
			log.Printf("Extracted filename: %s", filename)
		}

		if n.Kind() == ast.KindFencedCodeBlock {
			codeBlock := n.(*ast.FencedCodeBlock)
			if filename != "" {
				codeContent := extractCodeBlockContent(codeBlock, content)
				err := writeFile(filename, codeContent, tmpPrefix)
				if err != nil {
					log.Printf("Error writing to file: %s", err)
					return ast.WalkStop, err
				}
				log.Printf("Successfully wrote to file: %s", filename)
				filename = "" // Reset filename after writing the file
			}
		}

		return ast.WalkContinue, nil
	})

	if err != nil {
		log.Printf("Error walking through AST: %s", err)
		return fmt.Errorf("failed to walk through AST: %w", err)
	}

	log.Println("Finished parsing markdown file")
	return nil
}

func extractFilename(text string) string {
	// Extract the filename from the text in the format "**filename**:"
	filename := strings.TrimPrefix(text, "**")
	filename = strings.TrimSuffix(filename, "**:")
	return filename
}

func extractCodeBlockContent(codeBlock *ast.FencedCodeBlock, source []byte) string {
	var buf bytes.Buffer
	for i := 0; i < codeBlock.Lines().Len(); i++ {
		line := codeBlock.Lines().At(i)
		buf.Write(line.Value(source))
	}
	return buf.String()
}

func writeFile(filename, content string, tmpPrefix string) error {
	filename = tmpPrefix + filename  // Prepend the temporary directory prefix to the filename

	// Ensure the directory structure exists
	if err := os.MkdirAll(getDir(filename), os.ModePerm); err != nil {
		return err
	}
	return ioutil.WriteFile(filename, []byte(content), 0644)
}

func getDir(filename string) string {
	return filename[:strings.LastIndex(filename, "/")]
}
