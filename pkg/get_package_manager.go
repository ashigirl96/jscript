package pkg

import (
	"bufio"
	"errors"
	"fmt"
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
	stdpty, stdtty, _ := pty.Open()
	defer func(stdtty *os.File) {
		if err := stdtty.Close(); err != nil {
		}
	}(stdtty)
	cmd.Stdin = stdtty
	cmd.Stdout = stdtty
	errpty, errtty, _ := pty.Open()
	defer func(errtty *os.File) {
		if err := errtty.Close(); err != nil {
		}
	}(errtty)
	cmd.Stderr = errtty
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		return err
	}
	go func() {
		_, err := io.Copy(os.Stdout, stdpty)
		if err != nil {
			return
		}
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	go func() {
		_, err := io.Copy(os.Stderr, errpty)
		if err != nil {
			return
		}
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil

}
