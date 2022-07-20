package main

import (
	"github.com/attapon-th/gmf/cmd/gmf/command"
)

var (
	AppName   string
	Version   string
	Build     string
	DateBuild string
)

func init() {
	command.AppName = AppName
	command.Version = Version
	command.Build = Build
	command.DateBuild = DateBuild
}
func main() {
	command.Execute()
}
