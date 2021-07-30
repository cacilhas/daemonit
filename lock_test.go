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
		expected := fmt.Sprintf("/tmp/testRun.%s.lock", currentUser.Username)
		if filename := lockfile("testRun"); filename != expected {
			t.Fatalf("Expected %v, got %v", expected, filename)
		}
	})

	t.Run("lock", func(t *testing.T) {
		// Lock file
		var filename = lockfile("testRun")
		if err := lock("testRun"); err != nil {
			t.Fatalf("error locking daemon: %v", err)
		}
		if _, err := os.Lstat(filename); err != nil {
			t.Fatalf("error checking lock file: %v", err)
		}
		if err := lock("testRun"); err == nil {
			t.Fatalf("expected error, but nothing raised")
		}
		cleanupLock("testRun")
		if _, err := os.Lstat(filename); !os.IsNotExist(err) {
			t.Fatalf("expected file not to exist")
		}
	})
}
