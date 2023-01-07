/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Scripts = map[string]string

type PackageJson struct {
	Scripts Scripts `json:"scripts"`
}

func (p *PackageJson) String() string {
	var result []string
	for name, command := range p.Scripts {
		s := fmt.Sprintf("\x1b[32m%s\x1b[m:", name)
		script := fmt.Sprintf("%-24s%s", s, command)
		result = append(result, script)
	}
	return strings.Join(result, "\n")
}

var packageJson PackageJson

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "jscript",
	RunE: func(cmd *cobra.Command, args []string) error {
		packageJsonPath, err := cmd.Flags().GetString("package")
		if err != nil {
			return err
		}
		raw, err := os.ReadFile(packageJsonPath)
		if err != nil {
			return err
		}
		err = json.Unmarshal(raw, &packageJson)
		if err != nil {
			return err
		}
		fmt.Println(&packageJson)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("package", "p", "./package.json", "path of package.json")
}
