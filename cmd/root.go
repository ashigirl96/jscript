/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type PackageJson struct {
	Scripts map[string]string `json:"scripts"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "jscript",
	Run: func(cmd *cobra.Command, args []string) {
		packageJson, err := cmd.Flags().GetString("package")
		if err != nil {
			log.Fatal(err)
		}
		raw, err := os.ReadFile(packageJson)
		if err != nil {
			log.Fatal(err)
		}
		var result PackageJson
		err = json.Unmarshal([]byte(raw), &result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v", result)
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
	fmt.Println("HELLO!!")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("package", "p", "./package.json", "path of package.json")
}
