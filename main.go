package main
import (
"flag"
"fmt"
"os"
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
flag.Usage = func() {
fmt.Fprintf(os.Stderr, "Schema Drift Observatory (v%s)\n\n", Version)
fmt.Fprintf(os.Stderr, "Usage:\n")
fmt.Fprintf(os.Stderr, "  schema-drift-observatory --baseline <path> --candidate <path>\n\n")
fmt.Fprintf(os.Stderr, "Flags:\n")
flag.PrintDefaults()
}
flag.Parse()
if *versionFlag {
fmt.Printf("Schema Drift Observatory v%s\n", Version)
os.Exit(0)
}
if *baselineFlag == "" || *candidateFlag == "" {
fmt.Fprintln(os.Stderr, "Error: Both --baseline and --candidate flags are required.")
flag.Usage()
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
// ৬. OpenAPI Parsing
baseDoc, err := parser.ParseSpec(*baselineFlag)
if err != nil {
fmt.Fprintf(os.Stderr, "Baseline Error: %v\n", err)
os.Exit(2) // Invalid spec is a usage/input error
}
candDoc, err := parser.ParseSpec(*candidateFlag)
if err != nil {
fmt.Fprintf(os.Stderr, "Candidate Error: %v\n", err)
os.Exit(2)
}
fmt.Printf("Successfully parsed Baseline (Version: %s) and Candidate (Version: %s)\n", baseDoc.GetVersion(), candDoc.GetVersion())
os.Exit(0)
}
