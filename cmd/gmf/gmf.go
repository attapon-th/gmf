package main

import (
	"github.com/attapon-th/gmf/cmd/gmf/command"
)

var (
	AppName   string = "gmf"
	Version   string = "2.0.5"
	Build     string = "612ebe9dfcc2a57779af093f757007c2add35dce"
	DateBuild string = "2022-07-20T17:30:42+07:00"
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
