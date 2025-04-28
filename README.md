# Code-Decoder: AI-Powered Codebase Tutorial Generator

Code-Decoder transforms complex codebases into audience-targeted tutorials using AI. It can analyze GitHub repositories or local directories, identify core abstractions and interactions, and generate comprehensive, visualized documentation tailored to different audiences.

- [Code-Decoder: AI-Powered Codebase Tutorial Generator](#code-decoder-ai-powered-codebase-tutorial-generator)
  - [Features](#features)
  - [Installation](#installation)
    - [Prerequisites](#prerequisites)
    - [Installation Options](#installation-options)
      - [Option 1: Install from source](#option-1-install-from-source)
      - [Option 2: Install directly from the GitHub repo](#option-2-install-directly-from-the-github-repo)
      - [Option 3: Docker](#option-3-docker)
    - [Post-Installation Setup](#post-installation-setup)
  - [Usage Guide](#usage-guide)
    - [Detailed Command Documentation](#detailed-command-documentation)
      - [Analyze Command](#analyze-command)
      - [Generate Command](#generate-command)
      - [Test-LLM Command](#test-llm-command)
  - [Shell Completion](#shell-completion)
    - [Bash](#bash)
    - [Zsh](#zsh)
    - [Fish](#fish)
    - [PowerShell](#powershell)
  - [License](#license)


## Features

- Analyze GitHub repositories or local codebases
- Generate audience-targeted tutorials (beginner, developer, contributor)
- Create appropriate visualizations based on audience needs
- Support multiple output languages
- Modular LLM provider system with multiple options:
  - OpenAI and Anthropic (cloud-based)
  - Ollama and LM Studio (local, offline use)
- Save intermediate analysis for reuse
- Customize file inclusion/exclusion patterns
- Generate output in both Markdown and HTML formats

## Installation

### Prerequisites

- Go 1.24 or higher
- Make
- Git
- An API key for OpenAI or Anthropic (for cloud-based LLMs)
- Alternatively, Ollama or LM Studio for local LLM deployment

### Installation Options

#### Option 1: Install from source

1. Clone the repository:

   ```bash
   git clone https://github.com/ksylvan/code-decoder.git
   cd code-decoder
   ```

2. Install dependencies:

   ```bash
   make tidy
   ```

3. Build the application:

   ```bash
   make build
   ```

4. (Optional) Install the binary to your $GOPATH/bin:

   ```bash
   make install
   ```

#### Option 2: Install directly from the GitHub repo

   ```bash
   go install https://github.com/ksylvan/code-decoder/cmd/code-decoder
   ```

This should install the `code-decoder` binary in your  `$GOPATH/bin` directory (usually `~/go/bin`)

#### Option 3: Docker

TBD: Fill in the Docker build and docker run instructions.

### Post-Installation Setup

`code-decoder` loads its configuration from a `config.yaml` file in the current directory or `~/.config/code-decoder/`. You can override this by specifying the `--config` flag, which takes precedence over both locations.

1. Create a `config.yaml` file in your working directory:

   ```yaml
   llm:
      provider: "openai"  # Options: openai, anthropic, ollama, lmstudio
      api_key: ""  # Add your API key here for cloud providers
      model: "gpt-4"
      endpoint: ""  # Only needed for local providers

   defaults:
      output_dir: "./tutorials"
      language: "English"
      audience: "developer"
      include: ["*.go", "*.js", "*.py", "*.java", "*.rs", "*.c", "*.cpp", "*.h"]
      exclude: ["vendor/*", "node_modules/*", "*.test.js"]
      max_size: 1000000  # 1MB

   github:
      token: ""  # For private repositories
   ```

2. Set up your LLM provider:
   - For OpenAI: Get an API key from [OpenAI](https://platform.openai.com/api-keys)
   - For Anthropic: Get an API key from [Anthropic](https://console.anthropic.com/)
   - For Ollama: [Install Ollama](https://ollama.ai/) and run it locally
   - For LM Studio: [Install LM Studio](https://lmstudio.ai/) and run it locally

3. Test your LLM connection:

   ```bash
   code-decoder test-llm
   ```

## Usage Guide

Code-Decoder provides three main sub-commands:

1. `analyze`: Analyze a codebase and save the analysis to a file
2. `generate`: Generate tutorials from a codebase or saved analysis
3. `test-llm`: Test the connection to the LLM provider

### Detailed Command Documentation

#### Analyze Command

The `analyze` command processes a codebase and saves a structured analysis to a file.

```bash
code-decoder analyze [flags]
```

Required flags:

- `--dir` or `--repo`: Source code location (use exactly one)
- `--save-analysis`: File to save the analysis to

Optional flags:

- `--name`: Custom project name
- `--token`: GitHub token for private repositories
- `--include`: File patterns to include (comma-separated)
- `--exclude`: File patterns to exclude (comma-separated)
- `--max-size`: Maximum file size to include in bytes
- `--verbose`: Enable verbose output

Examples:

```bash
# Analyze a GitHub repository
code-decoder analyze --repo https://github.com/golang/go --save-analysis golang-analysis.json

# Analyze a local directory with custom filters
code-decoder analyze --dir ./my-project --name "My Project" --include="*.go,*.js" --exclude="test/*,vendor/*" --save-analysis my-analysis.json

# Analyze a private GitHub repository
code-decoder analyze --repo https://github.com/company/private-repo --token $GITHUB_TOKEN --save-analysis private-analysis.json
```

#### Generate Command

The `generate` command creates tutorials from a codebase or a saved analysis.

```bash
code-decoder generate [flags]
```

Required flags:

- Either `--load-analysis` or (`--dir` or `--repo`)

Optional flags:

- `--audience`: Target audience (beginner, developer, contributor)
- `--language`: Tutorial language (e.g., English, Chinese)
- `--output`: Directory to save generated tutorials
- `--format`: Output format (markdown, html)
- `--save-analysis`: Save the analysis to a file (if analyzing a codebase)
- `--provider`: Override the LLM provider
- `--verbose`: Enable verbose output

Examples:

```bash
# Generate a tutorial directly from a GitHub repository
code-decoder generate --repo https://github.com/user/project --audience beginner

# Generate a developer-focused tutorial from a saved analysis
code-decoder generate --load-analysis my-analysis.json --audience developer --output ./dev-docs

# Generate a tutorial in a different language and format
code-decoder generate --load-analysis my-analysis.json --audience contributor --language Chinese --format html --output ./zh-docs

# Generate a tutorial and save the analysis for later
code-decoder generate --dir ./my-project --save-analysis my-project.json --audience beginner
```

#### Test-LLM Command

The `test-llm` command verifies the connection to the configured LLM provider.

```bash
code-decoder test-llm [flags]
```

Optional flags:

- `--provider`: Override the LLM provider from config

Examples:

```bash
# Test the default LLM provider from config
code-decoder test-llm

# Test a specific provider
code-decoder test-llm --provider openai
```

## Shell Completion

`code-decoder` provides shell completion support for Bash, Zsh, Fish, and PowerShell.

To generate the completion script for your shell:

```bash
code-decoder completion [bash|zsh|fish|powershell]
```

Follow the instructions below for your specific shell:

### Bash

To load completions for the current session:

```bash
source <(code-decoder completion bash)
```

To load completions for each session, execute once:

```bash
# Linux:
sudo code-decoder completion bash > /etc/bash_completion.d/code-decoder

# macOS (using Homebrew):
code-decoder completion bash > "$(brew --prefix)/etc/bash_completion.d/code-decoder"
```

### Zsh

If shell completion is not already enabled in your environment, enable it by executing the following once:

```zsh
echo "autoload -U compinit; compinit" >> ~/.zshrc
```

To load completions for each session, execute once:

```zsh
# Ensure the completions directory exists (adjust path if necessary)
mkdir -p ~/.zsh/completion
code-decoder completion zsh > ~/.zsh/completion/_code-decoder
# Add the directory to your fpath in ~/.zshrc if it's not already there
# e.g., fpath=(~/.zsh/completion $fpath)
```

You will need to start a new shell for this setup to take effect.

Alternatively, if you use a framework like Oh My Zsh, you might place the completion file in its custom completions directory (e.g., `~/.oh-my-zsh/custom/plugins/code-decoder/` and add `code-decoder` to your plugins list).

### Fish

To load completions for the current session:

```fish
code-decoder completion fish | source
```

To load completions for each session, execute once:

```fish
code-decoder completion fish > ~/.config/fish/completions/code-decoder.fish
```

### PowerShell

To load completions for the current session:

```powershell
code-decoder completion powershell | Out-String | Invoke-Expression
```

To load completions for every new session, run:

```powershell
code-decoder completion powershell > code-decoder.ps1
```

Then, add the following line to your PowerShell profile (find its location by running `$profile`):

```powershell
. "<path-to-your-directory>\code-decoder.ps1"
```

## License

Copyright (c) 2025 Kayvan Sylvan

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
