package daemonit

func DaemonIt(callback func([]string) error, arg0 string, args []string) error {
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
