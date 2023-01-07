/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/ashigirl96/jscript/pkg"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "jscript",
	//PreRunE: func(cmd *cobra.Command, args []string) error {
	//	packageJsonPath, err := cmd.Flags().GetString("package")
	//	if err != nil {
	//		return err
	//	}
	//	return pkg.ReadPackageJson(packageJsonPath)
	//},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(&pkg.PackageJson)
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

//func init() {
//	// Cobra also supports local flags, which will only run
//	// when this action is called directly.
//	rootCmd.Flags().StringP("package", "p", "./package.json", "path of package.json")
//}
