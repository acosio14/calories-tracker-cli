package storage

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestGetOutputFilePath(t *testing.T) {
	got := GetOutputFilePath("/tmp/out", "adrian")
	want := filepath.Join("/tmp/out", "adrian.json")
	if got != want {
		t.Errorf("GetOutputFilePath = %q, want %q", got, want)
	}
}

func TestGetOutputFolderPath(t *testing.T) {
	got, err := GetOutputFolderPath()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasSuffix(got, "calories-tracker-output") {
		t.Errorf("folder path = %q, want it to end in calories-tracker-output", got)
	}
}
