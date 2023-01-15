/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/ashigirl96/jscript/pkg"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run one script that selected from package.json scripts",
	Args: func(cmd *cobra.Command, args []string) error {
		script := []string{args[0]}
		if err := cobra.OnlyValidArgs(cmd, script); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if pkg.PackageJsonDir.Dir != "" {
			return errors.New("cannot run this directory")
		}
		manager, err := pkg.GetPackageManager()
		if err != nil {
			return err
		}
		err = manager.Run(args...)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return nil
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if err := pkg.ReadPackageJson(); err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		if len(args) == 0 {
			commands := pkg.PackageJson.GetCommands()
			return commands, cobra.ShellCompDirectiveNoFileComp
		}
		return nil, cobra.ShellCompDirectiveFilterFileExt
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
