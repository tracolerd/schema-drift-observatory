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
