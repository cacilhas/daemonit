package daemonit

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

func lock(args []string) error {
	var err error
	var filename string

	if filename, err = lockfile(args); err != nil {
		return err
	}
	if err = checkFileExists(filename); err != nil {
		return err
	}
	pid := fmt.Sprintf("%d", os.Getpid())
	return ioutil.WriteFile(filename, []byte(pid), 0644)
}

func lockfile(args []string) (string, error) {
	var currentUser *user.User
	var err error
	processPath := strings.Split(args[0], "/")
	if currentUser, err = user.Current(); err != nil {
		return "", err
	}
	return fmt.Sprintf("/tmp/%s.%s.lock", processPath[len(processPath)-1], currentUser.Username), nil
}

func checkFileExists(filename string) error {
	if dat, err := ioutil.ReadFile(filename); err == nil {
		if pid, err := strconv.ParseInt(string(dat), 10, 32); err == nil {
			if process, err := os.FindProcess(int(pid)); err == nil {
				if err := process.Signal(syscall.Signal(0)); err == nil {
					return fmt.Errorf("process %d is running", pid)
				}
			}
		}
	} else if !os.IsNotExist(err) {
		return err
	}
	return nil
}

func cleanupLock(args []string) {
	if filename, err := lockfile(args); err == nil {
		os.Remove(filename)
	}
}
