package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersionCommandPrintsVersion(t *testing.T) {
	var out bytes.Buffer
	got := run([]string{"version"}, &out)
	if got != exitOK {
		t.Fatalf("version command exit = %d, want %d", got, exitOK)
	}
	if !strings.HasPrefix(out.String(), "vkc "+version+"\n") {
		t.Fatalf("version output = %q, want prefix %q", out.String(), "vkc "+version)
	}
}

func TestHelpCommandListsKnownCommands(t *testing.T) {
	var out bytes.Buffer
	if got := run([]string{"help"}, &out); got != exitOK {
		t.Fatalf("help command exit = %d, want %d", got, exitOK)
	}
	for _, want := range []string{"version", "help"} {
		if !strings.Contains(out.String(), want) {
			t.Errorf("help output %q does not mention %q", out.String(), want)
		}
	}
}

// Inverse cases: bad input must fail loudly with exit code 1 (usage error).

func TestUnknownCommandFailsWithUsageExit(t *testing.T) {
	var out bytes.Buffer
	got := run([]string{"definitely-not-a-command"}, &out)
	if got != exitUsage {
		t.Fatalf("unknown command exit = %d, want %d", got, exitUsage)
	}
	if out.String() != "" {
		t.Errorf("unknown-command error must go to stderr, stdout wrote %q", out.String())
	}
}

func TestNoArgumentsFailsWithUsageExit(t *testing.T) {
	var out bytes.Buffer
	if got := run(nil, &out); got != exitUsage {
		t.Fatalf("no-args exit = %d, want %d", got, exitUsage)
	}
}
