package ocl

import (
	"errors"
	"os/exec"
)

var ErrEmptyTarget = errors.New("empty target")
var ErrNoOpenerFound = errors.New("no suitable opener command found in PATH")

// Openers we’ll check for, in order of preference.
var openers = []string{
	"browser-exec",
	"xdg-open",
	"cmd.exe",
	"cygstart",
	"start",
	"open",
}

// Open is a smart open function that opens files, URLs, or other resources
// using the default application associated with the resource type on the
// current operating system.
func Open(target string) error {
	if target == "" {
		return ErrEmptyTarget
	}

	for _, opener := range openers {
		if _, err := exec.LookPath(opener); err == nil {
			return exec.Command(opener, target).Start()
		}
	}

	return ErrNoOpenerFound
}
