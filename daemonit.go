package daemonit

func DaemonIt(callback func([]string) error, args []string) error {
	daemon := true
	effectiveArgs := make([]string, 0, len(args))
	i := 0
	arg0 := args[0]
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
		fork(args)
		return nil
	} else {
		res := callback(effectiveArgs[:i])
		cleanupLock(arg0)
		return res
	}
}
