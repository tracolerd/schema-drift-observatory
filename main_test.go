package main

import (
	"strings"
	"testing"

	"github.com/tracolerd/schema-drift-observatory/internal/engine"
)

func TestValidateFormat(t *testing.T) {
	for _, format := range []string{"text", "json"} {
		if err := validateFormat(format); err != nil {
			t.Errorf("validateFormat(%q) returned unexpected error: %v", format, err)
		}
	}
	if err := validateFormat("yaml"); err == nil {
		t.Fatal("validateFormat(\"yaml\") expected an error")
	}
}

func TestParseFailOn(t *testing.T) {
	got, err := parseFailOn("breaking, potentially_breaking")
	if err != nil {
		t.Fatalf("parseFailOn returned unexpected error: %v", err)
	}
	want := []engine.Category{engine.CategoryBreaking, engine.CategoryPotentiallyBreaking}
	if len(got) != len(want) {
		t.Fatalf("got %d categories, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("category %d = %q, want %q", i, got[i], want[i])
		}
	}

	for _, invalid := range []string{"braking", "", "breaking,,non_breaking"} {
		if _, err := parseFailOn(invalid); err == nil {
			t.Errorf("parseFailOn(%q) expected an error", invalid)
		}
	}
	if _, err := parseFailOn("unknown"); err != nil && !strings.Contains(err.Error(), "unsupported") {
		t.Errorf("unknown category should produce a clear unsupported-category error, got %v", err)
	}
}
