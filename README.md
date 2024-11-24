# Markdown File Generator CLI

A command-line tool that generates files and directory structures from markdown files. This tool allows you to define your project structure and file contents in a markdown file and automatically generates the corresponding files and directories.

## Features

- 🔄 Generate files and directories from markdown specifications
- 📝 Support for multiple markdown header formats

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/will-wright-eng/parse
cd parse

# Initialize the project
make init

# Build the project
make build

# Optionally, install to your $GOPATH/bin
make install
```

### Prerequisites

- Go 1.21 or higher
- Make (for using the Makefile commands)

## Usage

### Basic Command

```bash
parse generate -i input.md -o output_dir
```

### Available Flags

- `-i, --input`: Input markdown file (required)
- `-o, --output`: Output directory (default: "./tmp")
- `-f, --force`: Force overwrite existing files
- `--strip-comments`: Strip comments from code blocks
- `--skip`: Patterns of files to skip (e.g., "*.tmp,*.bak")

### Markdown File Format

The tool supports multiple formats for specifying files in your markdown:

    ## path/to/file1.txt
    ```python
    def hello():
        print("Hello")
    ```

    file: path/to/file2.js
    ```javascript
    console.log('Hello');
    ```

    path/to/file3.go
    ```go
    package main

    func main() {
        println("Hello")
    }
    ```

### Project Structure

```bash
.
├── Makefile              # Build and development commands
├── README.md            # Project documentation
├── cmd/                 # Command-line interface
│   ├── generate.go      # Generate command implementation
│   ├── root.go         # Root command configuration
│   └── version.go      # Version command
├── internal/           # Internal packages
│   ├── config/        # Configuration handling
│   ├── generator/     # File generation logic
│   ├── logger/        # Logging utilities
│   ├── parser/        # Markdown parsing
│   └── version/       # Version information
└── main.go            # Application entry point
```

## Development

### Available Make Commands

```bash
Usage: make [command]

Commands:
  help             Display this help screen
  init             Project initialization
  add-pkgs         Add packages to go.mod
  run              Run the application
  run-generate     Run the generate command
  watch            Run the application with live reload
  build            Build the application
  build-linux      Build for Linux
  build-darwin     Build for macOS
  build-windows    Build for Windows
  test             Run tests
  test-coverage    Generate test coverage report
  benchmark        Run benchmarks
  lint             Run linters
  fmt              Format code
  vet              Run go vet
  deps             Install dependencies
  deps-update      Update dependencies
  clean            Clean build artifacts
  install          Install the application
  uninstall        Uninstall the application
  envs             Print environment variables
```

## Example file format


    config/settings.json

    ```json
    {
    "environment": "development",
    "port": 3000,
    "database": {
        "host": "localhost",
        "port": 5432
    }
    }
    ```

    src/main.js

    ```javascript
    console.log('Application starting...');
    ```


Running:

```bash
parse generate -i template.md -o myproject
```

Will create:

```
myproject/
├── config/
│   └── settings.json
└── src/
    └── main.js
```

---
Built with ❤️ using Go
