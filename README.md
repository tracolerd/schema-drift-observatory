# Schema Drift Observatory

**Detect breaking API contract changes before they reach consumers.**

Schema Drift Observatory is a planned API compatibility and contract-analysis tool for teams that need to compare API definitions across releases. It aims to turn specification differences into explainable, reviewable findings that developers can use locally or in continuous integration.

> **Project status:** Initial planning and repository setup. The comparison engine, dashboard, and CI integration are not implemented yet. Features below describe the intended design, not completed functionality.

## Why this project exists

An API can remain syntactically valid while becoming incompatible with existing clients. Removing a response field, changing a type, making an optional request field required, or removing an enum value can introduce a breaking change without an obvious failure during the provider's build.

Schema Drift Observatory is intended to compare a known baseline with a candidate specification, explain relevant changes, and produce machine-readable results that can be reviewed before release.

## Design goals

- **Explainable findings:** Every classification should link to a specific location and documented compatibility rule.
- **Directional analysis:** Compatibility is evaluated from the perspective of a defined consumer/provider relationship, not from a vague universal score.
- **Deterministic results:** The same inputs and rules should produce the same report.
- **Automation-friendly output:** CI should be able to consume a stable report and exit status.
- **Incremental scope:** Start with OpenAPI; add other contract formats only after the first engine is reliable.

## Planned capabilities

### Initial MVP

- Accept baseline and candidate OpenAPI specifications in JSON or YAML.
- Parse and validate both specifications, reporting useful errors for invalid inputs.
- Normalize relevant schema and operation details before comparison.
- Detect selected changes to paths, HTTP methods, parameters, request bodies, response schemas, required fields, types, and enum values.
- Classify findings as breaking, potentially breaking, or non-breaking using documented rules.
- Include evidence such as the affected operation, schema location, change summary, rule identifier, and suggested remediation.
- Export a JSON report and provide a predictable process exit code for CI usage.

### Later phases

- A Vue-based interface for inspecting changes and filtering findings.
- GitHub Actions examples and configurable release policies.
- GraphQL schema comparison.
- gRPC and Protocol Buffers compatibility checks.
- Historical comparisons and PostgreSQL-backed report storage, if history becomes a product requirement.

These later items are roadmap goals, not promises that every protocol will be supported in the first release.

## Proposed architecture

~~~mermaid
flowchart TD
    A[Baseline OpenAPI] --> C[Parse and Validate]
    B[Candidate OpenAPI] --> C
    C --> D[Normalization Layer]
    D --> E[Compatibility Rule Engine]
    E --> F[Structured Findings]
    F --> G[JSON Report and Exit Code]
    F --> H[Future Web Dashboard]
    G --> I[CI Policy Gate]
~~~

The Go comparison engine should remain usable independently of the web interface. Compatibility rules should be separated from parsing and reporting so that each rule can be tested with focused fixtures.

## Proposed technology

| Area | Planned choice | Rationale |
|---|---|---|
| Comparison engine | Go | A straightforward way to ship a fast, portable CLI |
| Contract format | OpenAPI 3.x | A useful, well-defined starting point for HTTP APIs |
| Web interface | Vue 3, TypeScript, Vite | A separate UI for reviewing detailed findings |
| Reports | JSON | Easy to inspect locally and consume in automation |
| CI | GitHub Actions | Demonstrate release-policy integration |
| Future persistence | PostgreSQL | Only needed if historical runs and shared workflows require storage |
| Future formats | GraphQL SDL, Protobuf | Separate adapters after OpenAPI rules stabilize |

## Compatibility rules and edge cases

A change should not be labelled breaking solely because its text differs. Rules must consider request versus response direction, required versus optional fields, schema constraints, nullable behavior where applicable, and whether a consumer may legally depend on the changed contract.

Initial test fixtures should cover at least:

- Removed paths and HTTP methods.
- A newly required request property.
- Removed response properties.
- Changed field types.
- Added optional response properties.
- Removed enum values.
- Parameter location or requiredness changes.
- Invalid specifications and unsupported constructs.

Each rule needs a documented rationale and tests. Where the impact cannot be established confidently, the engine should report uncertainty rather than invent certainty.

## Development roadmap

1. **Specification baseline:** Define normalized internal types, parser boundaries, CLI inputs, and representative fixtures.
2. **Comparison engine:** Implement the first documented OpenAPI rules with unit tests.
3. **Report contract:** Stabilize finding IDs, locations, severity categories, JSON output, and exit behavior.
4. **CLI hardening:** Improve validation messages, edge-case handling, and repeatable test coverage.
5. **Dashboard:** Build a Vue interface around the stable report contract.
6. **CI integration:** Add a sample workflow and configurable policy thresholds.
7. **Additional formats:** Evaluate GraphQL and Protobuf support against the same principles.

## Quality strategy

- Unit tests for each compatibility rule.
- Golden fixtures for representative baseline/candidate pairs.
- Negative tests for malformed input and unsupported definitions.
- Determinism tests for repeated comparisons.
- CLI integration tests for exit codes and report generation.
- Frontend tests once the dashboard exists.
- CI checks for formatting, linting, and tests once configured.

## Security and responsible handling

Specifications are untrusted input. The implementation should limit input size, handle parser failures safely, avoid executing specification content, and avoid leaking local filesystem paths or secrets in reports. Reports may reveal internal API structure; users should decide what can be uploaded or published.

## Repository status

This repository currently contains the project brief and license. No working CLI, API compatibility rules, dashboard, or automated test suite should be assumed to exist until code and verified test results are added.

## License

MIT. See [LICENSE](LICENSE).

## Features
- **Deterministic JSON Reports:** Ideal for CI/CD integration.
- **Strict Directional Checking:** Understands that adding a required request parameter is breaking, but adding an optional response property is not.
- **Configurable Exit Policies:** Fail your CI build only on the severity levels you care about.
- **Vue 3 Dashboard:** Includes a standalone SPA to visualize compatibility reports.

## Local Usage
`ash
# Build the binary
go build -o schema-drift-observatory main.go

# Run a comparison
./schema-drift-observatory --baseline base.yaml --candidate cand.yaml --fail-on breaking
`",
    ",
    
See examples/github-actions-usage.yml for a complete CI workflow.
