package daemonit

import (
	"os/exec"
	"time"
)

func fork(args []string) {
	command := args[0]
	args = append(args[1:], "--no-daemon")
	exec.Command(command, args...).Start()
	time.Sleep(200 * time.Millisecond)
}
