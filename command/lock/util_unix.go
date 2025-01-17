// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !windows
// +build !windows

package lock

import (
	"syscall"
)

// signalPid sends a sig signal to the process with process id pid.
func signalPid(pid int, sig syscall.Signal) error {
	return syscall.Kill(pid, sig)
}
