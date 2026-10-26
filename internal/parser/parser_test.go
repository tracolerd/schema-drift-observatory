package parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseSpecRejectsOversizedInput(t *testing.T) {
	t.Setenv(maxSpecBytesEnv, "8")
	path := filepath.Join(t.TempDir(), "large.yaml")
	if err := os.WriteFile(path, []byte("openapi: 3.0.0\ninfo: too big\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	_, err := ParseSpec(path)
	if err == nil || !strings.Contains(err.Error(), "maximum size of 8 bytes") {
		t.Fatalf("expected explicit size-limit error, got %v", err)
	}
}

func TestConfiguredMaxSpecBytesRejectsInvalidValues(t *testing.T) {
	for _, value := range []string{"0", "-1", "not-a-number", "104857601"} {
		t.Run(value, func(t *testing.T) {
			t.Setenv(maxSpecBytesEnv, value)
			if _, err := configuredMaxSpecBytes(); err == nil {
				t.Fatalf("expected invalid %s value %q to be rejected", maxSpecBytesEnv, value)
			}
		})
	}
}

func TestConfiguredMaxSpecBytesDefaultsToTenMiB(t *testing.T) {
	t.Setenv(maxSpecBytesEnv, "")
	got, err := configuredMaxSpecBytes()
	if err != nil {
		t.Fatal(err)
	}
	if got != 10*1024*1024 {
		t.Fatalf("got %d bytes, want %d", got, 10*1024*1024)
	}
}
