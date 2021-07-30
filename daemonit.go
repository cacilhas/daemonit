package daemonit

import (
	"os"
)

func DaemonIt(callback func([]string) error, args []string) error {
	if args == nil {
		args = os.Args
	}
	daemon := true
	effectiveArgs := make([]string, 0, len(args))
	i := 0
	for _, arg := range args {
		if arg == "--no-daemon" {
			daemon = false
		} else {
			effectiveArgs[i] = arg
			i++
		}
	}
	if daemon {
		if err := lock(args); err != nil {
			return err
		}
		fork(args)
		return nil
	} else {
		res := callback(effectiveArgs[:i])
		cleanupLock(args)
		return res
	}
}
