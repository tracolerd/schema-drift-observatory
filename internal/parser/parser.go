package parser

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/pb33f/libopenapi"
)

const (
	defaultMaxSpecBytes      int64 = 10 * 1024 * 1024
	maxConfigurableSpecBytes int64 = 100 * 1024 * 1024
	maxSpecBytesEnv                = "SCHEMA_DRIFT_MAX_SPEC_BYTES"
)

func configuredMaxSpecBytes() (int64, error) {
	value := os.Getenv(maxSpecBytesEnv)
	if value == "" {
		return defaultMaxSpecBytes, nil
	}
	limit, err := strconv.ParseInt(value, 10, 64)
	if err != nil || limit <= 0 || limit > maxConfigurableSpecBytes {
		return 0, fmt.Errorf("%s must be an integer between 1 and %d", maxSpecBytesEnv, maxConfigurableSpecBytes)
	}
	return limit, nil
}

// ParseSpec reads and parses an OpenAPI document while bounding input memory.
// The default limit is 10 MiB and can be raised up to 100 MiB with
// SCHEMA_DRIFT_MAX_SPEC_BYTES.
func ParseSpec(filePath string) (libopenapi.Document, error) {
	maxBytes, err := configuredMaxSpecBytes()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to inspect specification file: %w", err)
	}
	if info.Size() > maxBytes {
		return nil, fmt.Errorf("specification exceeds the maximum size of %d bytes (%s)", maxBytes, maxSpecBytesEnv)
	}

	// Read one byte beyond the configured limit to remain safe if the file
	// changes after Stat or reports an unexpected size.
	data, err := io.ReadAll(io.LimitReader(file, maxBytes+1))
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	if int64(len(data)) > maxBytes {
		return nil, fmt.Errorf("specification exceeds the maximum size of %d bytes (%s)", maxBytes, maxSpecBytesEnv)
	}

	doc, err := libopenapi.NewDocument(data)
	if err != nil {
		return nil, fmt.Errorf("invalid OpenAPI specification: %w", err)
	}
	return doc, nil
}
