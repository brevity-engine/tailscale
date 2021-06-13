// Copyright (c) 2021 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dns

import (
	"os/exec"
	"strings"
	"syscall"
)

// We access WSL2 file systems via wsl.exe instead of \\wsl$\ because
// the netpath appears to operate as the standard user, not root.

func wslFileExists(distro string, fileName string) (bool, error) {
	cmd := exec.Command("wsl.exe", "-u", "root", "-d", distro, "/bin/test", fileName)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	if err == nil {
		return true, nil
	}
	if ee, _ := err.(*exec.ExitError); ee != nil {
		if ee.ExitCode() == 1 {
			return false, nil
		}
	}
	return false, err
}

func wslReadFile(distro, fileName string) (string, error) {
	cmd := exec.Command("wsl.exe", "-u", "root", "-d", distro, "/bin/cat", fileName)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	b, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}

func wslWriteFile(distro, fileName, contents string) error {
	panic("TODO")
}
