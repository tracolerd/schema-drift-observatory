package engine
import (
"fmt"
"github.com/tracolerd/schema-drift-observatory/internal/model"
)
// rulePathRemoved checks if an endpoint in the baseline is missing in the candidate.
func rulePathRemoved(baseline, candidate *model.NormalizedSpec, result *Result) {
for key, baseEndpoint := range baseline.Endpoints {
if _, exists := candidate.Endpoints[key]; !exists {
addFinding(result, Finding{
RuleID:      "SDO-001",
Category:    CategoryBreaking,
Method:      baseEndpoint.Method,
Path:        baseEndpoint.Path,
Location:    fmt.Sprintf("paths.%s.%s", baseEndpoint.Path, baseEndpoint.Method),
Explanation: fmt.Sprintf("Operation %s %s was removed in candidate spec.", baseEndpoint.Method, baseEndpoint.Path),
Impact:      "Existing clients calling this endpoint will receive HTTP 404/405 errors.",
Remediation: "Restore the operation or mark it as deprecated before removal in a major version.",
})
}
}
}
