package daemonit

import (
	"os"
)

func DaemonIt(callback func([]string) error, args []string) error {
	daemon := true
	argsLen := len(args)
	var arg0 string
	if argsLen == 0 {
		arg0, _ = os.Executable()
	} else {
		arg0 = args[0]
	}
	effectiveArgs := make([]string, argsLen)
	i := 0
	for _, arg := range args {
		if arg == noDaemonParam {
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
		if args[0] == arg0 {
			fork(arg0, args[1:])
		} else {
			fork(arg0, args)
		}
		return nil

	} else {
		res := callback(effectiveArgs[:i])
		cleanupLock(arg0)
		return res
	}
}
