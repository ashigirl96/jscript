/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/ashigirl96/jscript/cmd"
	"github.com/ashigirl96/jscript/pkg"
)

func main() {
	err := pkg.ReadPackageJson("./package.json")
	if err != nil {
		os.Exit(1)
	}
	cmd.Execute()
}
