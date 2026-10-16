package engine
import (
"fmt"
"github.com/tracolerd/schema-drift-observatory/internal/model"
)
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
// SDO-002: Request parameter added as required
func ruleParamAddedRequired(baseline, candidate *model.NormalizedSpec, result *Result) {
for epKey, candEp := range candidate.Endpoints {
baseEp, exists := baseline.Endpoints[epKey]
for paramKey, candParam := range candEp.Parameters {
if candParam.Required {
if !exists || baseEp.Parameters[paramKey] == nil {
addFinding(result, Finding{
RuleID:      "SDO-002",
Category:    CategoryBreaking,
Method:      candEp.Method,
Path:        candEp.Path,
Location:    fmt.Sprintf("paths.%s.%s.parameters.%s", candEp.Path, candEp.Method, candParam.Name),
Explanation: fmt.Sprintf("A new required %s parameter '%s' was added.", candParam.In, candParam.Name),
Impact:      "Clients not sending this new required parameter will receive HTTP 400 Bad Request.",
Remediation: "Make the parameter optional, or bump the major API version.",
})
}
}
}
}
}
// SDO-003: Request parameter removed or made required
func ruleParamChanged(baseline, candidate *model.NormalizedSpec, result *Result) {
for epKey, baseEp := range baseline.Endpoints {
candEp, candExists := candidate.Endpoints[epKey]
if !candExists {
continue // Already handled by SDO-001
}
for paramKey, baseParam := range baseEp.Parameters {
candParam, paramExists := candEp.Parameters[paramKey]
if !paramExists {
addFinding(result, Finding{
RuleID:      "SDO-003",
Category:    CategoryPotentiallyBreaking,
Method:      baseEp.Method,
Path:        baseEp.Path,
Location:    fmt.Sprintf("paths.%s.%s.parameters.%s", baseEp.Path, baseEp.Method, baseParam.Name),
Explanation: fmt.Sprintf("The %s parameter '%s' was removed.", baseParam.In, baseParam.Name),
Impact:      "Clients sending this parameter might face issues if the server strictly validates unknown parameters.",
Remediation: "Ensure the server safely ignores unknown parameters, or deprecate it first.",
})
} else if !baseParam.Required && candParam.Required {
addFinding(result, Finding{
RuleID:      "SDO-003",
Category:    CategoryBreaking,
Method:      baseEp.Method,
Path:        baseEp.Path,
Location:    fmt.Sprintf("paths.%s.%s.parameters.%s", baseEp.Path, baseEp.Method, candParam.Name),
Explanation: fmt.Sprintf("The %s parameter '%s' was changed from optional to required.", baseParam.In, baseParam.Name),
Impact:      "Clients not sending this parameter will now receive HTTP 400 Bad Request.",
Remediation: "Keep the parameter optional to maintain backward compatibility.",
})
}
}
}
}
