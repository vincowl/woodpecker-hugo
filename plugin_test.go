package main

import (
	"testing"

	"github.com/pkg/errors"
)

func TestVersionEqual(t *testing.T) {
	want := true
	if got := versionsEqual("1.0", "1.0"); want != got {
		t.Errorf("want: %t, got: %t", want, got)
	}

	want = false
	if got := versionsEqual("1.5", "1.0"); want != got {
		t.Errorf("want: %t, got: %t", want, got)
	}
}

func argsEqual(want []string, got []string) error {
	for i := range want {
		if want[i] != got[i] {
			return errors.Errorf("Arguments do not match, want: %s, got: %s", want[i], got[i])
		}
	}
	return nil
}
