package daemonit

import (
	"os"
)

func DaemonIt(arg0 string, callback func([]string) error, args []string) error {
	if args == nil {
		args = os.Args
	}
	if arg0 == "" {
		arg0, _ = os.Executable()
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
		if err := lock(arg0); err != nil {
			return err
		}
		fork(arg0, args)
		return nil
	} else {
		res := callback(effectiveArgs[:i])
		cleanupLock(arg0)
		return res
	}
}
