package configuration

import (
	"strings"
	"testing"
)

func TestLoad_errorForNonExistingFile(t *testing.T) {
	_, err := Load("path/to/non-existing.yaml")
	if err == nil {
		t.Errorf("expected non-nil error")
	}

	cerr, ok := err.(ConfigurationError)
	if !ok {
		t.Errorf("expected ConfigurationError")
	}

	expectedMsg := "open path/to/non-existing.yaml: no such file or directory"
	if cerr.Error() != expectedMsg {
		t.Errorf("unexpected message: %s", err.Error())
	}
}

func TestLoad_errorInvalidYamlContent(t *testing.T) {
	_, err := Load("testdata/invalid.yml")
	if err == nil {
		t.Errorf("expected non-nil error")
	}

	cerr, ok := err.(ConfigurationError)
	if !ok {
		t.Errorf("expected ConfigurationError")
	}

	if !strings.Contains(cerr.Error(), "yaml: unmarshal errors:") {
		t.Errorf("unexpected message: %s", err.Error())
	}
}

func TestLoad_success(t *testing.T) {
	_, err := Load("testdata/valid.yml")
	if err != nil {
		t.Errorf("failed with error: %v", err)
	}
}
