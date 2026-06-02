//go:build windows

package user

import "syscall"

func helmInstallSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{}
}
