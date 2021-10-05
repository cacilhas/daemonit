package daemonit

import (
	"os/exec"
	"time"
)

var noDaemonParam string = "--no-daemon"

func fork(command string, args []string) {
	exec.Command(command, append(args, noDaemonParam)...).Start()
	time.Sleep(200 * time.Millisecond)
}
