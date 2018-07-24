package fileutil

import "testing"

func TestOpen(t *testing.T) {
	got := Open()
	if got != 0 {
		t.Errorf("Expected 0")
	}
}