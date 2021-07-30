package daemonit

import (
	"os/exec"
	"time"
)

func fork(arg0 string, args []string) {
	args = append(args, "--no-daemon")
	exec.Command(arg0, args...).Start()
	time.Sleep(200 * time.Millisecond)
}
