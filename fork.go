package daemonit

import (
	"os/exec"
	"time"
)

func fork(args []string) {
	args = append(args, "--no-daemon")
	exec.Command(args[0], args[1:]...).Start()
	time.Sleep(200 * time.Millisecond)
}
