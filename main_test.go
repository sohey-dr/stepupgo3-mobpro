package main

import (
	"strings"
	"testing"
)

func TestRunCommandSuccess(t *testing.T) {
	results, err := runCommand("github.com/gostaticanalysis/skeleton/v2")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(results) == 0 {
		t.Fatal("failed test")
	}
}

func TestRunTipOneSuccess(t *testing.T) {
	versions, err := runTipOne("github.com/gostaticanalysis/skeleton/v2")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(versions) == 0 {
		t.Fatal("failed test")
	}
	for _, version := range versions {
		if !strings.HasPrefix(version, "v2") {
			t.Fatal("failed test")
		}
	}
}
