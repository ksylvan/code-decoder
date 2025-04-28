<!--
Copyright (c) 2025 Kayvan Sylvan

This source code is licensed under the MIT license found in the
LICENSE file in the root directory of this source tree.
-->

# Code-Decoder: AI-Powered Codebase Tutorial Generator

Requirements for a Go-based AI Powered Codebase Tutorial Generator.

## Overview

Your task is to design a detailed project architecture for a tool that transforms complex codebases into audience-targeted tutorials using AI. The project must be implemented in Go and should provide a seamless way to switch between various LLM providers, including frontier models from providers like OpenAI and Anthropic, as well as locally hosted models via Ollama and LM Studio. The goal is to replicate the functionality described in the provided README, focusing on crawling GitHub repositories or local directories, analyzing code structures, and generating detailed, visualized tutorials for users with different levels of expertise.

## Project Goals

- Develop an AI agent that analyzes GitHub repositories or local codebases to create audience-targeted tutorials.
- Identify core abstractions and interactions within codebases to explain how components work together.
- Generate tiered documentation suitable for different user personas (beginners, developers, contributors).
- Create appropriate visualizations based on the target audience's needs.
- Support multiple languages for tutorial output, with ability to translate source code comments.
- Ensure the tool is user-friendly with a simple setup and execution process.

## Functional Requirements

### Input and Source Management

1. **Repository Analysis**: Clone and scan GitHub repositories (public and private).
2. **Local Directory Analysis**: Scan local code directories.
3. **Custom Filtering**: Allow users to specify file inclusion/exclusion patterns.
4. **Size Limiting**: Skip files exceeding a configurable size limit.

### Code Analysis and Knowledge Extraction

1. **Project Overview Extraction**: Identify purpose, technologies, and main functionalities.
2. **Component Analysis**: Identify key modules, classes, and functions.
3. **Relationship Mapping**: Analyze dependencies between components.
4. **Workflow Analysis**: Identify key execution flows and data processes.
5. **Analysis Storage**: Save intermediate analysis results for reuse.

### Tutorial Generation

1. **Audience-Targeted Content**: Generate different content tiers:
   - Beginner: Focus on core concepts and simple examples
   - Developer: Focus on API usage and integration points
   - Contributor: Focus on internal architecture and extension points
2. **Multilingual Support**: Generate tutorials in multiple languages.
3. **Code Examples**: Extract and explain representative code examples.
4. **Progressive Complexity**: Structure tutorials to build from simple to complex concepts.
5. **Navigation Aids**: Include tables of contents, cross-references, etc.

### Visualization

1. **Architecture Diagrams**: Generate system-level architecture visualizations.
2. **Component Diagrams**: Create visualizations of key components and relationships.
3. **Sequence Diagrams**: Illustrate important process flows.
4. **Audience-Specific Detail**: Adjust visualization complexity based on audience.

### Output Formats

1. **Markdown Generation**: Create well-structured Markdown documents.
2. **HTML Generation**: Generate HTML documentation with navigation and styling.
3. **Embeddable Visualizations**: Include rendered diagrams in output files.

### AI Integration

1. **Multiple LLM Provider Support**: Interface with:
   - OpenAI API (gpt-4, etc.)
   - Anthropic API (Claude models)
   - Ollama (local models)
   - LM Studio (local models)
2. **Provider Fallback**: Allow cascading through providers if one fails.
3. **Context Management**: Efficiently manage context windows for large codebases.
4. **Prompt Engineering**: Develop specialized prompts for different analysis phases.

## Non-Functional Requirements

### Performance

1. **Efficient Analysis**: Process medium-sized codebases (50-100 files) within 10-15 minutes.
2. **Parallel Processing**: Analyze multiple files concurrently where appropriate.
3. **Resource Management**: Avoid excessive memory usage when processing large codebases.

### Reliability

1. **Error Handling**: Gracefully handle API failures, rate limits, and connectivity issues.
2. **Result Validation**: Verify AI-generated content for consistency and accuracy.
3. **Session Resumption**: Allow for resuming interrupted analysis sessions.

### Usability

1. **Simple Setup**: Minimal configuration required for basic usage.
2. **Clear Documentation**: Comprehensive guides for installation and usage.
3. **Helpful Error Messages**: Provide actionable feedback for common errors.
4. **Progress Indicators**: Show status during lengthy operations.

### Security

1. **API Key Management**: Secure storage and handling of API credentials.
2. **Data Privacy**: Ensure code analysis happens locally when requested.
3. **Network Security**: Secure communication with external APIs.

### Extensibility

1. **Plugin Architecture**: Allow for adding new LLM providers easily.
2. **Custom Templates**: Support user-provided templates for outputs.
3. **Output Format Extensions**: Framework for adding new output formats.
4. **Language Support**: Easily add support for new tutorial languages.

## Technical Requirements

### Language and Libraries

1. **Go Implementation**: Primary codebase in Go 1.21+.
2. **Dependency Management**: Use Go modules for dependency management.
3. **CLI Framework**: Utilize cobra for command-line interface.
4. **HTTP Clients**: Standard library or well-maintained clients for API communication.
5. **Template Engine**: Use Go's template system for output generation.
6. **Visualization Tools**: Integration with Mermaid.js or similar for diagrams.

### Architecture

1. **Modular Design**: Clear separation of concerns between components.
2. **Interface-Driven**: Use interfaces to abstract provider implementations.
3. **Configuration Management**: Support for config files, env vars, and CLI args.
4. **Testing Strategy**: Unit tests for core functionality, integration tests for end-to-end flows.

### Development Practices

1. **Code Quality**: Follow Go best practices for code organization and style.
2. **Documentation**: GoDoc comments for all exported functions and types.
3. **Error Handling**: Consistent approach to error propagation and reporting.
4. **Logging**: Structured logging with configurable verbosity levels.
5. **Versioning**: Semantic versioning for releases.

## Deliverables

1. **Executable Application**: Cross-platform Go binary.
2. **Source Code**: Well-structured and commented codebase.
3. **Documentation**: Installation, usage, and configuration guides.
4. **Sample Outputs**: Example tutorials generated from common open-source projects.
5. **Test Suite**: Comprehensive tests covering core functionality.

## Success Criteria

1. Successfully analyze and generate tutorials for at least 3 different types of codebases (e.g., web service, CLI tool, library).
2. Generate distinguishably different outputs for the three audience types.
3. Support all specified LLM providers with a unified interface.
4. Process a codebase of at least 100 files with reasonable performance.
5. Generate visually appealing and informative diagrams as part of the tutorials.
