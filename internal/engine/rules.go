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
func ruleParamChanged(baseline, candidate *model.NormalizedSpec, result *Result) {
for epKey, baseEp := range baseline.Endpoints {
candEp, candExists := candidate.Endpoints[epKey]
if !candExists { continue }
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
// Helpers for Schema Required checks
func isRequired(propName string, reqList []string) bool {
for _, r := range reqList {
if r == propName { return true }
}
return false
}
// SDO-004: Request body property newly required
func ruleRequestBodyChanged(baseline, candidate *model.NormalizedSpec, result *Result) {
for epKey, candEp := range candidate.Endpoints {
baseEp, exists := baseline.Endpoints[epKey]
if !exists || candEp.RequestBody == nil || candEp.RequestBody.Properties == nil { continue }
for propName := range candEp.RequestBody.Properties {
isCandReq := isRequired(propName, candEp.RequestBody.Required)
if baseEp.RequestBody == nil || baseEp.RequestBody.Properties == nil || baseEp.RequestBody.Properties[propName] == nil {
// New Property added
if isCandReq {
addFinding(result, Finding{
RuleID:      "SDO-004", Category: CategoryBreaking,
Method: candEp.Method, Path: candEp.Path,
Location:    fmt.Sprintf("paths.%s.%s.requestBody.%s", candEp.Path, candEp.Method, propName),
Explanation: fmt.Sprintf("A new required property '%s' was added to the request body.", propName),
Impact:      "Clients not sending this property will receive HTTP 400.",
Remediation: "Make the new request property optional.",
})
}
} else {
// Existing property changed to required
isBaseReq := isRequired(propName, baseEp.RequestBody.Required)
if !isBaseReq && isCandReq {
addFinding(result, Finding{
RuleID:      "SDO-004", Category: CategoryBreaking,
Method: candEp.Method, Path: candEp.Path,
Location:    fmt.Sprintf("paths.%s.%s.requestBody.%s", candEp.Path, candEp.Method, propName),
Explanation: fmt.Sprintf("Request body property '%s' was changed from optional to required.", propName),
Impact:      "Clients not sending this property will receive HTTP 400.",
Remediation: "Keep the property optional.",
})
}
}
}
}
}
// SDO-005 & 006: Response property removed or optional added
func ruleResponseBodyChanged(baseline, candidate *model.NormalizedSpec, result *Result) {
for epKey, baseEp := range baseline.Endpoints {
candEp, exists := candidate.Endpoints[epKey]
if !exists || baseEp.Responses["200"] == nil || candEp.Responses["200"] == nil { continue }
baseSchema := baseEp.Responses["200"].Content
candSchema := candEp.Responses["200"].Content
if baseSchema == nil || baseSchema.Properties == nil { continue }
// Check for removed properties (SDO-005)
for propName := range baseSchema.Properties {
if candSchema == nil || candSchema.Properties == nil || candSchema.Properties[propName] == nil {
addFinding(result, Finding{
RuleID:      "SDO-005", Category: CategoryBreaking,
Method: baseEp.Method, Path: baseEp.Path,
Location:    fmt.Sprintf("paths.%s.%s.responses.200.%s", baseEp.Path, baseEp.Method, propName),
Explanation: fmt.Sprintf("Response property '%s' was removed.", propName),
Impact:      "Clients expecting this property may fail to parse the response.",
Remediation: "Restore the property or deprecate it over a major version.",
})
}
}
// Check for newly added optional properties (SDO-006)
if candSchema != nil && candSchema.Properties != nil {
for propName := range candSchema.Properties {
if baseSchema.Properties[propName] == nil {
addFinding(result, Finding{
RuleID:      "SDO-006", Category: CategoryNonBreaking,
Method: candEp.Method, Path: candEp.Path,
Location:    fmt.Sprintf("paths.%s.%s.responses.200.%s", candEp.Path, candEp.Method, propName),
Explanation: fmt.Sprintf("A new property '%s' was added to the response.", propName),
Impact:      "Safe for existing clients that ignore unknown properties.",
Remediation: "None required.",
})
}
}
}
}
}
