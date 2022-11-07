package my_docker

import (
	"os/exec"
	"syscall"
	"testing"
)

func TestNs(t *testing.T) {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:
	}
}
