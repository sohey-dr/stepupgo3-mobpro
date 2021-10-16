package main

import (
	"strings"
	"testing"
)

func TestRunSuccess(t *testing.T) {
	results, err := run("github.com/gostaticanalysis/skeleton/v2")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(results) == 0 {
		t.Fatal("resultsが空です")
	}
}

func TestTipOneSuccess(t *testing.T) {
	versions, err := tipOne("github.com/gostaticanalysis/skeleton/v2")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(versions) == 0 {
		t.Fatal("versionsが空です")
	}
	for _, version := range versions {
		if !strings.HasPrefix(version, "v2") {
			t.Fatal("failed test")
		}
	}
}

func TestTipTwoSuccess(t *testing.T) {
	path := "github.com/gostaticanalysis/skeleton/v2"
	filePath, err := tipTwo(path)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if filePath == "" {
		t.Fatal("filePathが空です")
	}
	if !strings.Contains(filePath, path) {
		t.Fatal("対象のモジュールを検索できてません")
	}
}

func TestTipThreeSuccess(t *testing.T) {
	output, _ := tipThree("./...")
	if !strings.Contains(output, "{") {
		t.Fatal("対象のモジュールをgo vetできていません")
	}
}
