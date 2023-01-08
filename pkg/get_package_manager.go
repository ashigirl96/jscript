package pkg

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
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
func (m *Manager) Run(command string) error {
	cmd := exec.Command(m.String(), "run", command)
	stdout, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		return err
	}
	outBuffer := make([]byte, 100)
	for {
		if _, err := stdout.Read(outBuffer); err != nil {
			return err
		}
		r := bufio.NewReader(stdout)
		line, _, _ := r.ReadLine()
		fmt.Println(string(line))
	}
}
