# Code-Decoder Project Roadmap

This roadmap tracks the implementation and testing progress for the features outlined in the [Product Requirements Document (PRD)](PRD.md).

| Phase / Requirement                                         | Code | Tests |
| :---------------------------------------------------        | :--- | :---: |
| **Phase 1: Minimal CLI Skeleton & Configuration**           |      |       |
| 1.1 Implement Cobra CLI (`analyze`, `generate`, `test-llm`) | ✅   | ☑️    |
| 1.2 Provide help/usage output for commands                  | ✅   | ☑️    |
| 1.3 Implement config loading (flags, env, files)            | ✅   | ☑️    |
| 1.4 Validate config & display errors                        | ✅   | ☑️    |
| 1.5 Implement shell completion command                      | ✅   | ☑️    |
||||
| **Phase 2: Source Management & File Listing**               |      |       |
| 2.1 Implement `SourceManager` (dir/repo validation)         |      |       |
| 2.2 List files based on include/exclude/size                |      |       |
| 2.3 Integrate file listing into `analyze` command           |      |       |
||||
| **Phase 3: Code Parsing & Basic Analysis**                  |      |       |
| 3.1 Implement language detection                            |      |       |
| 3.2 Implement basic `Parser` interface (e.g., Go)           |      |       |
| 3.3 Parse files (path, lang, imports, decls)                |      |       |
| 3.4 Output summary of parsed files via CLI                  |      |       |
||||
| **Phase 4: LLM Client Integration & Knowledge Extraction**  |      |       |
| 4.1 Implement `Provider` interface (e.g., OpenAI)           |      |       |
| 4.2 Implement `test-llm` command                            |      |       |
| 4.3 Implement `KnowledgeExtractor`                          |      |       |
| 4.4 Integrate knowledge extraction into `analyze`           |      |       |
||||
| **Phase 5: Tutorial Generation & Rendering**                |      |       |
| 5.1 Implement `ContentGenerator` (audience-based)           |      |       |
| 5.2 Implement `TemplateEngine` (Markdown)                   |      |       |
| 5.3 Add `generate` command for Markdown output              |      |       |
||||
| **Phase 6: Visualization & Advanced Output**                |      |       |
| 6.1 Implement `Visualizer` (arch, component, diagrams)      |      |       |
| 6.2 Integrate diagrams into Markdown output                 |      |       |
| 6.3 Extend `TemplateEngine` for HTML output                 |      |       |
||||
| **Phase 7: Error Handling, Testing, and Extensibility**     |      |       |
| 7.1 Implement domain-specific errors/wrapping               |      |       |
| 7.2 Add unit and integration tests                          |      |       |
| 7.3 Provide mocking interfaces (LLM, source manager)        |      |       |
| 7.4 Document/support adding new providers/parsers           |      |       |

Legend:

- ⏳ In Progress
- ✅ Done (in the context of tests, this means we have automated tests)
- ☑️ Manually tested (no automated tests)
- ❌ Blocked
