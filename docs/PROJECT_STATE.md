# Project State: Schema Drift Observatory
## Current Phase
Phase 0: Architecture Foundation & Repository Setup
## Technical Stack Overview
Our primary engine is built using Go as a standalone CLI application to ensure fast execution and static binary compilation. For OpenAPI parsing and validation, we will rely on a maintained Go library (to be chosen in Phase 1). Future expansions will include a UI dashboard built with Vue 3, TypeScript, and Vite, keeping the architecture strictly modular without any database, microservices, or AI/LLM dependencies for the MVP.
## Completed Milestones
- [x] Environment Setup (Go 1.27.0 installed, Repo Cloned)
- [x] Base Documentation Created (PROJECT_STATE, DECISIONS, SPEC_SUBSET_AND_CONTRACT)
## Upcoming Milestones
- [ ] Phase 1: Go Module, CLI Skeleton & Input Validation
- [ ] Phase 2: Normalized Model & Core Rule Engine with Fixtures
- [ ] Phase 3: Deterministic JSON Report, Terminal Summary & Exit Policy
- [ ] Phase 4: Edge cases, Property/Golden Tests, Linter/CI Setup
- [ ] Phase 5: Vue 3 + TypeScript Dashboard
- [ ] Phase 6: GitHub Actions policy example
- [ ] Phase 7: Evaluate GraphQL/Protobuf adapters
## Known Blockers & Risks
1. **Windows vs Unix Path & Line Endings:** Windows CRLF vs LF and path separators could break golden fixtures. Mitigation: Normalize paths and line endings in reports and tests.
2. **OpenAPI 3.0 vs 3.1 Differences:** Nullable types and JSON schema versions differ. Mitigation: Use a strict normalized internal model before rule execution.
3. **Complex Polymorphism:** Deeply nested oneOf/nyOf or external $refs can cause panics. Mitigation: Identify unsupported constructs early and flag them explicitly as undetermined.
