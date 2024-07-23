/*
Copyright © 2024 shynome <shynome@gmail.com>
*/
package main

import "github.com/vscode-lcode/hub/v2/cmd"

var Version = "dev"

func main() {
	cmd.Execute(Version)
}
