<!--
Copyright (c) 2025 Kayvan Sylvan

This source code is licensed under the MIT license found in the
LICENSE file in the root directory of this source tree.
-->

# Product Requirements Document (PRD) for Code-Decoder

## Phase 1: Minimal CLI Skeleton & Configuration

**Objective:**
Establish a runnable CLI application with basic command structure and configuration management, enabling initial user interaction and configuration validation.

**Requirements:**

1. Implement a CLI entry point using Cobra with the following subcommands:
    - `analyze`
    - `generate`
    - `test-llm`
2. Provide help and usage output for each command.
3. Implement configuration loading with precedence:
    - Command-line flags
    - Environment variables
    - Config file in current directory
    - Config file in user’s home directory
4. Validate configuration and display errors for missing or invalid settings.

**Testability:**

- CLI launches and displays help for each command.
- Running each subcommand with `--help` shows usage information.
- Configuration is loaded correctly from each source, with precedence respected.
- Invalid or missing configuration results in clear error messages.

**Dependencies:**
None (foundational phase).

---

## Phase 2: Source Management & File Listing

**Objective:**
Enable the CLI to accept a local directory or repository, validate it, and list source files for analysis.

**Requirements:**

1. Implement `SourceManager` to:
    - Accept a local directory path or repository URL.
    - Validate the source (existence, accessibility).
    - List files based on include/exclude patterns and size limits.
2. Integrate file listing into the `analyze` command, outputting the list of files to the user.

**Testability:**

- Providing a valid directory or repository lists files as expected.
- Invalid sources (nonexistent path, bad URL) produce clear errors.
- Include/exclude patterns and size limits are respected.

**Dependencies:**
Phase 1 (CLI and configuration must be in place).

---

## Phase 3: Code Parsing & Basic Analysis

**Objective:**
Parse source files to extract structural information and prepare for higher-level analysis.

**Requirements:**

1. Implement language detection for files.
2. Implement a basic `Parser` interface for at least one language (e.g., Go).
3. Parse files to extract:
    - File path
    - Language
    - Imports
    - Declarations (functions, types, etc.)
4. Output a summary of parsed files via the CLI.

**Testability:**

- Files are parsed and their structure is output in a summary table or JSON.
- Language detection is correct for supported files.
- Errors in parsing are reported clearly.

**Dependencies:**
Phase 2 (file listing and source management).

---

## Phase 4: LLM Client Integration & Knowledge Extraction

**Objective:**
Enable the system to interact with an LLM provider and extract high-level project knowledge.

**Requirements:**

1. Implement the `Provider` interface for at least one LLM (e.g., OpenAI).
2. Add the ability to test LLM connectivity via the `test-llm` command.
3. Implement `KnowledgeExtractor` to:
    - Extract project overview
    - Extract component relationships
    - Extract workflows
4. Integrate knowledge extraction into the `analyze` command, outputting results.

**Testability:**

- LLM connection test passes/fails as expected.
- Knowledge extraction produces summaries, relationships, and workflows for a sample project.
- Errors from the LLM are handled gracefully.

**Dependencies:**
Phase 3 (parsed code structure required for knowledge extraction).

---

## Phase 5: Tutorial Generation & Rendering

**Objective:**
Generate audience-targeted tutorials and render them in Markdown.

**Requirements:**

1. Implement `ContentGenerator` to create tutorial content based on analysis and audience.
2. Implement `TemplateEngine` to render tutorials to Markdown.
3. Add a `generate` command that produces a Markdown tutorial from analysis results.

**Testability:**

- Running `generate` produces a Markdown file with a title, introduction, table of contents, sections, and conclusion.
- Content is tailored to the specified audience.
- Output is well-formed Markdown.

**Dependencies:**
Phase 4 (analysis and knowledge extraction must be available).

---

## Phase 6: Visualization & Advanced Output

**Objective:**
Enhance tutorials with architecture and component diagrams, and support additional output formats.

**Requirements:**

1. Implement `Visualizer` to generate:
    - Architecture diagrams
    - Component diagrams
    - Sequence diagrams
2. Integrate diagrams into the generated Markdown.
3. Extend `TemplateEngine` to support HTML output.

**Testability:**

- Tutorials include generated diagrams in Markdown and HTML.
- Diagrams accurately reflect the codebase structure and workflows.
- HTML output is valid and visually consistent.

**Dependencies:**
Phase 5 (tutorial generation and rendering).

---

## Phase 7: Error Handling, Testing, and Extensibility

**Objective:**
Ensure robust error handling, comprehensive testing, and support for new languages and LLM providers.

**Requirements:**

1. Implement domain-specific error types and error wrapping.
2. Add unit and integration tests for all major components.
3. Provide mocking interfaces for LLM and source management.
4. Document and support the addition of new LLM providers and language parsers.

**Testability:**

- Errors are user-friendly and traceable.
- All components have passing unit and integration tests.
- New providers/parsers can be added following documented steps and are detected by the system.

**Dependencies:**
All previous phases (finalization and polish).
