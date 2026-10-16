package engine
import (
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
// Execute Rules
rulePathRemoved(baseline, candidate, result)
ruleParamAddedRequired(baseline, candidate, result)
ruleParamChanged(baseline, candidate, result)
// Update compatibility flag
if result.Counts[CategoryBreaking] > 0 {
result.Compatible = false
}
return result
}
func addFinding(res *Result, f Finding) {
res.Findings = append(res.Findings, f)
res.Counts[f.Category]++
}
