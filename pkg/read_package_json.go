package pkg

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Scripts = map[string]string

type IPackageJson struct {
	Scripts Scripts `json:"scripts"`
}

func (p *IPackageJson) String() string {
	commands := make([]string, 0, len(PackageJson.Scripts))
	for name, command := range p.Scripts {
		s := fmt.Sprintf("\x1b[32m%s\x1b[m:", name)
		command := fmt.Sprintf("%-24s%s", s, command)
		commands = append(commands, command)
	}
	return strings.Join(commands, "\n")
}

var PackageJson IPackageJson

func (p *IPackageJson) GetCommands() []string {
	commands := make([]string, 0, len(PackageJson.Scripts))
	for name, command := range PackageJson.Scripts {
		commands = append(commands, fmt.Sprintf("%s\t%s", name, command))
	}
	return commands
}

func ReadPackageJson(path string) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, &PackageJson)
	if err != nil {
		return err
	}
	return nil
}
