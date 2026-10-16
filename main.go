package main
import (
"encoding/json"
"flag"
"fmt"
"os"
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
// ৭. Evaluate Rules
result := engine.CompareSpecs(baseSpec, candSpec)
// ৮. Output JSON Report (Temporary for testing)
reportBytes, _ := json.MarshalIndent(result, "", "  ")
fmt.Println(string(reportBytes))
if !result.Compatible {
os.Exit(1) // Policy violation exit code
}
os.Exit(0)
}
