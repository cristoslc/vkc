// Package main implements vkc, a dynamic command-line client for Vikunja.
//
// vkc discovers its command surface at runtime from the OpenAPI spec each
// Vikunja instance serves at /api/v1/docs.json (ADR-0002, ADR-0004). This
// file is the v0 skeleton: version, help, and the usage-error exit contract.
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// version is the CLI's semantic version; wired to release tags later.
const version = "0.0.1"

// Exit codes are the automation contract (see docs/USER-EXPERIENCE.md):
// 0 success, 1 usage error, 2 transport failure, 3 auth failure.
const (
	exitOK    = 0
	exitUsage = 1
)

// errUsage marks a usage error; callers map it to exit code 1.
var errUsage = errors.New("usage error")

// run executes vkc with argv (without av[0]) against w for user-facing output.
// It returns the process exit code. Errors are reported to stderr by the caller.
func run(argv []string, w io.Writer) int {
	if len(argv) == 0 {
		fmt.Fprintln(os.Stderr, "usage: vkc <command> [args]")
		fmt.Fprintln(os.Stderr, "run 'vkc help' for available commands")
		return exitUsage
	}
	switch argv[0] {
	case "version":
		fmt.Fprintln(w, "vkc", version)
		return exitOK
	case "help":
		fmt.Fprintf(w, "vkc %s — dynamic Vikunja command-line client (scaffold)\n", version)
		fmt.Fprintln(w, "commands: version, help")
		return exitOK
	default:
		fmt.Fprintf(os.Stderr, "vkc: unknown command %q\n", argv[0])
		fmt.Fprintln(os.Stderr, "run 'vkc help' for available commands")
		return exitUsage
	}
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout))
}
