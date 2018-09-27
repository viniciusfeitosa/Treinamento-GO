package models_test

import (
	"testing"
)

func TestDadosDoPai(t *testing.T) {
	expected := "algum valor"

	f := funcaoQueFazAlgumaCoisa()

	if f != expected {
		t.Errorf("Expected: %s, Result: %s", expected, f)
	}
}
