/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/ashigirl96/jscript/cmd"
	"github.com/ashigirl96/jscript/pkg"
	"os"
)

func main() {
	err := pkg.ReadPackageJson("./package.json")
	if err != nil {
		os.Exit(1)
	}
	cmd.Execute()
}
