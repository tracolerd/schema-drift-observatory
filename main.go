package main
import (
"encoding/json"
"flag"
"fmt"
"os"
"strings"
"github.com/tracolerd/schema-drift-observatory/internal/engine"
"github.com/tracolerd/schema-drift-observatory/internal/parser"
)
const Version = "0.1.0"
func checkFileExists(path string) error {
info, err := os.Stat(path)
if os.IsNotExist(err) {
return fmt.Errorf("file not found: %s", path)
}
if err != nil {
return fmt.Errorf("error accessing file %s: %v", path, err)
}
if info.Size() == 0 {
return fmt.Errorf("file is empty: %s", path)
}
return nil
}
func main() {
versionFlag := flag.Bool("version", false, "Print tool version")
baselineFlag := flag.String("baseline", "", "Path to baseline OpenAPI spec (JSON/YAML)")
candidateFlag := flag.String("candidate", "", "Path to candidate OpenAPI spec (JSON/YAML)")
formatFlag := flag.String("format", "text", "Output format: text or json")
failOnFlag := flag.String("fail-on", "breaking", "Comma-separated categories to fail the CI build (e.g., breaking,potentially_breaking)")
flag.Parse()
if *versionFlag {
fmt.Printf("Schema Drift Observatory v%s\n", Version)
os.Exit(0)
}
if *baselineFlag == "" || *candidateFlag == "" {
fmt.Fprintln(os.Stderr, "Error: Both --baseline and --candidate flags are required.")
os.Exit(2)
}
if err := checkFileExists(*baselineFlag); err != nil {
fmt.Fprintf(os.Stderr, "Baseline Error: %v\n", err)
os.Exit(2)
}
if err := checkFileExists(*candidateFlag); err != nil {
fmt.Fprintf(os.Stderr, "Candidate Error: %v\n", err)
os.Exit(2)
}
baseDoc, err := parser.ParseSpec(*baselineFlag)
if err != nil {
fmt.Fprintf(os.Stderr, "Baseline Error: %v\n", err)
os.Exit(2)
}
candDoc, err := parser.ParseSpec(*candidateFlag)
if err != nil {
fmt.Fprintf(os.Stderr, "Candidate Error: %v\n", err)
os.Exit(2)
}
baseSpec, err := parser.ExtractSpec(baseDoc)
if err != nil {
fmt.Fprintf(os.Stderr, "Error extracting baseline: %v\n", err)
os.Exit(2)
}
candSpec, err := parser.ExtractSpec(candDoc)
if err != nil {
fmt.Fprintf(os.Stderr, "Error extracting candidate: %v\n", err)
os.Exit(2)
}
result := engine.CompareSpecs(baseSpec, candSpec)
// Format Output
if *formatFlag == "json" {
reportBytes, _ := json.MarshalIndent(result, "", "  ")
fmt.Println(string(reportBytes))
} else {
// Text Summary Output
fmt.Printf("Schema Drift Observatory Comparison Report\n")
fmt.Printf("----------------------------------------\n")
fmt.Printf("Baseline:  %s\n", *baselineFlag)
fmt.Printf("Candidate: %s\n", *candidateFlag)
fmt.Printf("\nSummary:\n")
fmt.Printf("  Breaking:             %d\n", result.Counts[engine.CategoryBreaking])
fmt.Printf("  Potentially Breaking: %d\n", result.Counts[engine.CategoryPotentiallyBreaking])
fmt.Printf("  Non-Breaking:         %d\n", result.Counts[engine.CategoryNonBreaking])
if len(result.Findings) > 0 {
fmt.Printf("\nFindings:\n")
for i, f := range result.Findings {
fmt.Printf("  %d. [%s] %s %s - %s\n", i+1, strings.ToUpper(string(f.Category)), f.Method, f.Path, f.Explanation)
fmt.Printf("     Rule: %s | Location: %s\n", f.RuleID, f.Location)
}
} else {
fmt.Printf("\nNo compatibility issues found!\n")
}
}
// Exit Code Policy Evaluation
failOnCategories := strings.Split(*failOnFlag, ",")
hasViolation := false
for _, cat := range failOnCategories {
c := engine.Category(strings.TrimSpace(cat))
if result.Counts[c] > 0 {
hasViolation = true
break
}
}
if hasViolation {
os.Exit(1)
}
os.Exit(0)
}
