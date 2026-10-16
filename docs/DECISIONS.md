# Architecture Decisions
## 2026-10-16: Core Architecture and Engine
**Context:** Need a fast, standalone API compatibility analysis tool.
**Decision:** Build the first engine as a Go CLI application.
**Rationale:** Go provides static binaries, fast execution, excellent CLI libraries, and strong ecosystem support for OpenAPI parsing. 
## 2026-10-16: Scope Constraint (MVP)
**Context:** OpenAPI spec parsing can become infinitely complex (e.g., recursive $ref, large polymorphism).
**Decision:** Restrict MVP to a documented subset of OpenAPI 3.0/3.1. Unsupported constructs will be explicitly marked as undetermined rather than ignored or falsely evaluated. No databases, no network interception, no microservices.
## 2026-10-16: Category Semantics
**Decision:** All findings will be classified into four explicit categories:
1. reaking: Breaks existing clients (e.g., required request parameter added).
2. potentially_breaking: May break clients depending on implementation (e.g., enum value removed).
3. 
on_breaking: Safe for existing clients (e.g., new optional property added).
4. undetermined: Spec constructs that are outside the supported MVP scope.

## 2026-10-16: Phase 7 - GraphQL & Protobuf Evaluation (MVP Wrap-up)
**Context:** The roadmap required evaluating GraphQL/Protobuf adapters after OpenAPI stabilization.
**Decision:** Deferred to a future v2.0 release. OpenAPI relies on REST semantics (Paths, Methods, HTTP Status codes), whereas GraphQL uses a single endpoint with query ASTs, and gRPC relies on strict Protobuf contracts. Integrating them requires completely separate parser engines and a different internal normalized model. MVP scope is successfully completed with OpenAPI 3.x support.
