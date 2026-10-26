# Prioritized Engineering Backlog

This is a planning backlog based on the current documented CLI/dashboard scope. Each item should be validated against the implementation before coding; no item should be treated as a confirmed defect until reproduced.

## Candidate issues

1. **Reject unsupported CLI format values** — Return a clear validation error for unknown `--format` values instead of silently falling back to text.
2. **Validate every `--fail-on` category** — Reject typos and unsupported categories so a misspelled CI policy cannot weaken release gating.
3. **Compare schema property types and enum values** — Report compatible/incompatible changes to existing properties with explicit rule IDs and tests.
4. **Handle success responses beyond 200 JSON** — Detect removed success statuses and media-type changes instead of only inspecting `200 application/json`.
5. **Emit an undetermined finding for unsupported constructs** — Make incomplete analysis visible rather than allowing missing findings to imply compatibility.
6. **Apply an input-size limit** — Bound specification reads and return an actionable error when a file exceeds the configured limit.
7. **Fix and test dashboard sample-report counts** — Ensure summary totals are derived from, or verified against, the findings array.
8. **Build and test the dashboard in CI** — Install frontend dependencies reproducibly and run its production build as part of CI.
9. **Align Go toolchain versions between CI and go.mod** — Choose an explicit supported version and verify local/CI parity.
10. **Make the PowerShell build script fail fast and run tests** — Stop after any failed command and include `go test ./...` before reporting success.

## Prioritization

Start with items 1–6 because incorrect compatibility reports or weakened CI policies can mislead consumers. Then fix fixture integrity and verification (7–10). For every rule change, add a minimal baseline/candidate fixture pair, tests for the expected result, and a nearby non-breaking or unsupported case.
