# Project State: Schema Drift Observatory
## Current Phase
Phase 6: GitHub Actions policy example & Documentation
## Technical Stack Overview
Our primary engine is built using Go as a standalone CLI application to ensure fast execution and static binary compilation. For OpenAPI parsing and validation, we will rely on a maintained Go library (`pb33f/libopenapi`). Future expansions will include a UI dashboard built with Vue 3, TypeScript, and Vite.
## Completed Milestones
- [x] Environment Setup (Go 1.27.0 installed, Repo Cloned)
- [x] Base Documentation Created (PROJECT_STATE, DECISIONS, SPEC_SUBSET_AND_CONTRACT)
- [x] Phase 1: Go Module, CLI Skeleton & Input Validation
- [x] Phase 2: Normalized Model & Core Rule Engine with Fixtures (SDO-001 to SDO-006)
- [x] Phase 3: Deterministic JSON Report, Terminal Summary & Exit Policy
- [x] Phase 4: Edge cases, Property/Golden Tests, Linter/CI Setup
- [x] Phase 5: Vue 3 + TypeScript Dashboard
## Upcoming Milestones
- [ ] Phase 6: GitHub Actions policy example & Documentation
- [ ] Phase 7: Evaluate GraphQL/Protobuf adapters
