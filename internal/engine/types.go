package engine
// Category defines the severity of a finding.
type Category string
const (
CategoryBreaking            Category = "breaking"
CategoryPotentiallyBreaking Category = "potentially_breaking"
CategoryNonBreaking         Category = "non_breaking"
CategoryUndetermined        Category = "undetermined"
)
// Finding represents a single compatibility issue or change detected.
type Finding struct {
RuleID      string   `json:"rule_id"`
Category    Category `json:"category"`
Method      string   `json:"method,omitempty"`
Path        string   `json:"path,omitempty"`
Location    string   `json:"location,omitempty"`
Explanation string   `json:"explanation"`
Impact      string   `json:"impact"`
Remediation string   `json:"remediation"`
}
// Result holds the final outcome of a comparison.
type Result struct {
Compatible bool             `json:"compatible"`
Counts     map[Category]int `json:"counts"`
Findings   []Finding        `json:"findings"`
}
