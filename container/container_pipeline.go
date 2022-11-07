package container

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

func NewPipe() (*os.File, *os.File, error) {
	read, write, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}
	return read, write, nil
}

func NewParentProcess(tty bool) (*exec.Cmd, *os.File) {
	readPipe, writePipe, err := NewPipe()
	if err != nil {
		log.Errorf("new pipe error %v", err)
		return nil, nil
	}
	cmd := exec.Command("/proc/self/exe", "init")
	cmd.SysProcAttr = &syscall.SysProcAttr{Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC}
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	//ExtraFiles属性表示会外带着这个文件句柄去创建子进程。一个进程有三个默认文件描述符:标准输入，标准输出，标准错误。进程已创建就会默认带着这三个，这个外带的fd(readPipe)会成为第四个。
	// /proc/self/fd中出现四个: 0->/dev/pts/5 1->/dev/pts/5 2->/dev/pts/5 3->/proc/20765/fd
	cmd.ExtraFiles = []*os.File{readPipe}
	return cmd, writePipe
}

func readUserCommand() []string {
	file, _ := os.OpenFile(uintptr(3))
	ioutil.ReadAll(file)
}
