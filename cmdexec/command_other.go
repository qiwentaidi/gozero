//go:build !windows

package cmdexec

import "os/exec"

func prepareCommand(cmd *exec.Cmd) {}
