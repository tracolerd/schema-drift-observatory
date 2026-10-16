package engine
import (
"github.com/tracolerd/schema-drift-observatory/internal/model"
)
// CompareSpecs compares a baseline spec against a candidate spec and returns a Result.
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
// Rule: SDO-001 (Path/Operation Removed)
rulePathRemoved(baseline, candidate, result)
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
