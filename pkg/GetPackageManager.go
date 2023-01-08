package pkg

import (
	"errors"
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
	if _, err := os.Stat("./package-lock.json"); err == nil {
		return Npm, nil
	}
	if _, err := os.Stat("./yarn.lock"); err == nil {
		return Yarn, nil
	}
	if _, err := os.Stat("./pnpm-lock.yaml"); err == nil {
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

func (m *Manager) Run(command string) ([]byte, error) {
	return exec.Command(m.String(), "run", command).Output()
}