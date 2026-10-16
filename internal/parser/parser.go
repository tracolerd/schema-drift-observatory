package parser
import (
"fmt"
"os"
"github.com/pb33f/libopenapi"
)
// ParseSpec reads and parses an OpenAPI document.
func ParseSpec(filePath string) (libopenapi.Document, error) {
data, err := os.ReadFile(filePath)
if err != nil {
return nil, fmt.Errorf("failed to read file: %w", err)
}
doc, err := libopenapi.NewDocument(data)
if err != nil {
return nil, fmt.Errorf("invalid OpenAPI specification: %w", err)
}
return doc, nil
}
