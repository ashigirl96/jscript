/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ashigirl96/jscript/pkg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "jscript",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := pkg.ReadPackageJson(); err != nil {
			return err
		}
		fmt.Println(&pkg.PackageJson)
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
