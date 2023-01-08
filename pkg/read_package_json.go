package pkg

import (
	"encoding/json"
	"fmt"
	"os"
	path2 "path"
	"strings"
)

type Scripts = map[string]string

type IPackageJson struct {
	Scripts Scripts `json:"scripts"`
}

type IPackageJsonDir struct{ Dir string }

var PackageJson IPackageJson
var PackageJsonDir IPackageJsonDir

func (p *IPackageJson) String() string {
	commands := make([]string, 0, len(PackageJson.Scripts))
	for name, command := range p.Scripts {
		s := fmt.Sprintf("\x1b[32m%s\x1b[m:", name)
		command := fmt.Sprintf("%-20s\t%s", s, command)
		commands = append(commands, command)
	}
	return strings.Join(commands, "\n")
}

func (p *IPackageJson) GetCommands() []string {
	commands := make([]string, 0, len(PackageJson.Scripts))
	for name, command := range PackageJson.Scripts {
		commands = append(commands, fmt.Sprintf("%s\t%s", name, command))
	}
	return commands
}

func ReadPackageJson() error {
	path := PackageJsonDir.Path("package.json")
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

func init() {
	PackageJsonDir = IPackageJsonDir{
		os.Getenv("PACKAGE_DIR"),
	}
}

func (d *IPackageJsonDir) Path(name string) string {
	return path2.Join(d.Dir, name)

}
