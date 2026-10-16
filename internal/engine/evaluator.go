package engine
import (
"sort"
"github.com/tracolerd/schema-drift-observatory/internal/model"
)
func CompareSpecs(baseline, candidate *model.NormalizedSpec) *Result {
result := &Result{
Compatible: true,
Counts: map[Category]int{
CategoryBreaking:            0,
CategoryPotentiallyBreaking: 0,
CategoryNonBreaking:         0,
CategoryUndetermined:        0,
},
Findings: []Finding{},
}
rulePathRemoved(baseline, candidate, result)
ruleParamAddedRequired(baseline, candidate, result)
ruleParamChanged(baseline, candidate, result)
ruleRequestBodyChanged(baseline, candidate, result)
ruleResponseBodyChanged(baseline, candidate, result)
if result.Counts[CategoryBreaking] > 0 {
result.Compatible = false
}
// Make JSON output deterministic by sorting findings
sort.SliceStable(result.Findings, func(i, j int) bool {
// Sort by RuleID, then Method, then Path, then Location
if result.Findings[i].RuleID != result.Findings[j].RuleID {
return result.Findings[i].RuleID < result.Findings[j].RuleID
}
if result.Findings[i].Method != result.Findings[j].Method {
return result.Findings[i].Method < result.Findings[j].Method
}
if result.Findings[i].Path != result.Findings[j].Path {
return result.Findings[i].Path < result.Findings[j].Path
}
return result.Findings[i].Location < result.Findings[j].Location
})
return result
}
func addFinding(res *Result, f Finding) {
res.Findings = append(res.Findings, f)
res.Counts[f.Category]++
}
