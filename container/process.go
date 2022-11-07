package container

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"os"
	"os/exec"
	"syscall"
)

func NewParentProcess(tty bool, command string) *exec.Cmd {
	args := []string{"init", command} //这个init就是之前定义的initCommand命令行！
	//当前进程调用自己进行init初始化操作  /proc/self/exe init ${command}对进程进行初始化
	cmd := exec.Command("/proc/self/exe", args...)
	//clone参数fork出新进程，使用namespace隔离新创建的进程和外部环境
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWNET,
	}
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd
}

// init进程如果先于用户进程执行的话，pid=1的进程就变成了init，非常不合理且不能退出，所以要用execve系统调用
// 调用kernel的int execve(const char *filename,char *const argv[],char *const envp[]),用于执行当前filename对应的程序。会覆盖调当前进程镜像数据和堆栈，包括pid都会被将要运行的进程覆盖掉。
// 这样就可以把init进程替换掉，让容器内第一个程序变成指定的进程
func RunContainerInitProcess(command string, args []string) error {
	log.Infof("command %s", command)
	//使用mount挂载proc文件系统，MS_NOEXEC本文件系统中不允许运行其他程序 MS_NOSUID本系统中运行程序的时候不允许setuserid或setgroupid MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(syscall.MS_NOEXEC|syscall.MS_NOSUID|syscall.MS_NODEV), "")
	argv := []string{command}
	if err := unix.Exec(command, argv, os.Environ()); err != nil {
		log.Errorf(err.Error())
	}
	return nil
}
