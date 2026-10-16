# Specification Subset & Output Contract
## Supported OpenAPI Subset (MVP)
- **Formats:** JSON and YAML
- **Versions:** OpenAPI 3.0.x and 3.1.x (basic support)
- **Validated Components:** Paths, HTTP Operations (GET, POST, PUT, PATCH, DELETE), Parameters (query, header, path), Request Body schemas, Response schemas (2xx), required fields, primitive data types, and enum values.
- **Unsupported (MVP):** External $ref, deep llOf/nyOf/oneOf resolution, callbacks, webhooks, and header casing nuances. These will return an undetermined finding.
## CLI Exit Code Policy
- 0: Success, comparison completed, no policy violations (e.g., no breaking changes if --fail-on=breaking).
- 1: Success, comparison completed, policy violation detected.
- 2: Usage error, invalid JSON/YAML, file not found, or unsupported OpenAPI major version.
