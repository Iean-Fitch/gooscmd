package oscmd

import (
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

var (
	osCmds = make(map[int]*exec.Cmd)
)

func StartCmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	fmt.Printf("cmd:%s, arg:%v \n", name, arg)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		log.Printf("Launch '%s' error:%v\n", name, err)
		return err
	}

	pid := cmd.Process.Pid
	osCmds[pid] = cmd

	return nil
}

func KillCmd(Pid int) {

	if cmd, ok := osCmds[Pid]; ok {
		err := syscall.Kill(cmd.Process.Pid, 9)
		if err != nil {
			log.Printf("Call syscall return error: %v\n", err)
			return
		}
		cmd.Wait()
	}
}

func KillAll() {

	for _, cmd := range osCmds {
		err := syscall.Kill(cmd.Process.Pid, 9)
		if err != nil {
			log.Printf("Call syscall return error: %v\n", err)
			continue
		}
		cmd.Wait()
	}
}
