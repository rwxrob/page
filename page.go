package page

import (
	"fmt"
	"os"
	"os/exec"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/to"
)

// FindPager returns a full path to a pager binary if it can find one on
// the system:
//
//     * $PAGER
//     * pager (in PATH)
//
// If neither is found returns an empty string.
func FindPager() string {
	path := os.Getenv(`PAGER`)
	if path == "" {
		path, _ = exec.LookPath(path)
	}
	if path == "" {
		path, _ = exec.LookPath(`pager`)
	}
	return path
}

// File looks up the system pager and passes the first argument to
// it.
func File(path string) error {
	pager := FindPager()
	if pager == "" {
		return fmt.Errorf(`failed to find pager`)
	}
	return Z.Exec(pager, path)
}

// Page pipes the buf as input to the system pager. Anything that
// to.String accepts can be passed.
func This[T any](buf T) error {
	oin := os.Stdin
	defer func() { os.Stdin = oin }()
	pager := FindPager()
	if pager == "" {
		return fmt.Errorf(`failed to find pager`)
	}
	f, err := os.CreateTemp("", `page-*`)
	if err != nil {
		return err
	}
	_, err = f.WriteString(to.String(buf))
	defer f.Close()
	defer os.Remove(f.Name())
	if err != nil {
		return err
	}
	return Z.Exec(pager, f.Name())
}
