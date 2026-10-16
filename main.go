package main
import (
"flag"
"fmt"
"os"
)
const Version = "0.1.0"
// checkFileExists verifies if a file exists and is not empty.
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
// ৫. File Validation: Check if files exist and are not empty
if err := checkFileExists(*baselineFlag); err != nil {
fmt.Fprintf(os.Stderr, "Baseline Error: %v\n", err)
os.Exit(2)
}
if err := checkFileExists(*candidateFlag); err != nil {
fmt.Fprintf(os.Stderr, "Candidate Error: %v\n", err)
os.Exit(2)
}
fmt.Println("Files successfully validated.")
os.Exit(0)
}
