package daemonit

import (
	"fmt"
	"os"
	"os/user"
	"testing"
)

func TestLockfile(t *testing.T) {
	t.Run("lockfile", func(t *testing.T) {
		var currentUser *user.User
		var err error
		if currentUser, err = user.Current(); err != nil {
			panic(err)
		}
		args := []string{"testRun", "lockfile"}
		expected := fmt.Sprintf("/tmp/testRun.%s.lock", currentUser.Username)
		if filename, _ := lockfile(args); filename != expected {
			t.Fatalf("Expected %v, got %v", expected, filename)
		}
	})

	t.Run("lock", func(t *testing.T) {
		// Lock file
		args := []string{"testRun", "lock"}
		var filename, _ = lockfile(args)
		if err := lock(args); err != nil {
			t.Fatalf("error locking daemon: %v", err)
		}
		if _, err := os.Lstat(filename); err != nil {
			t.Fatalf("error checking lock file: %v", err)
		}
		if err := lock(args); err == nil {
			t.Fatalf("expected error, but nothing raised")
		}
		cleanupLock(args)
		if _, err := os.Lstat(filename); !os.IsNotExist(err) {
			t.Fatalf("expected file not to exist")
		}
	})
}
