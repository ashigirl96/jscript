package pkg

import (
	"errors"
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

type Manager int

const (
	Npm Manager = iota
	Yarn
	Pnpm
	Other
)

func GetPackageManager() (Manager, error) {
	npmLock := PackageJsonDir.Path("package-lock.json")
	if _, err := os.Stat(npmLock); err == nil {
		return Npm, nil
	}
	yarnLock := PackageJsonDir.Path("yarn.lock")
	if _, err := os.Stat(yarnLock); err == nil {
		return Yarn, nil
	}
	pnpmLock := PackageJsonDir.Path("pnpm-lock.yaml")
	if _, err := os.Stat(pnpmLock); err == nil {
		return Pnpm, nil
	}
	return Other, errors.New("never support that manager")
}

func (m *Manager) String() string {
	switch *m {
	case Npm:
		return "npm"
	case Yarn:
		return "yarn"
	case Pnpm:
		return "pnpm"
	default:
		return ""
	}
}

// Run 親プロセスから切り出したい
func (m *Manager) Run(command ...string) error {
	cmd := exec.Command(m.String(), command...)
	stdPty, stdTty, _ := pty.Open()
	defer func(stdTty *os.File) {
		if err := stdTty.Close(); err != nil {
		}
	}(stdTty)
	cmd.Stdin = os.Stdin
	// cmd.Stdoutへの出力をstyTtyにする
	cmd.Stdout = stdTty
	errPty, errTty, _ := pty.Open()
	defer func(errTty *os.File) {
		if err := errTty.Close(); err != nil {
		}
	}(errTty)
	cmd.Stderr = errTty
	if err := cmd.Start(); err != nil {
		return err
	}
	go func() {
		_, err := io.Copy(os.Stdout, stdPty)
		if err != nil {
			return
		}
	}()
	go func() {
		_, err := io.Copy(os.Stderr, errPty)
		if err != nil {
			return
		}
	}()
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil

}
