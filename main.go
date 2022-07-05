package main

import (
	"github.com/attapon-th/gmf/services/checker"
	"github.com/attapon-th/phuslulogger"
	"github.com/phuslu/log"
)

func main() {
	phuslulogger.SetDefaultlogger()
	if err := checker.CheckGoVersion(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
	if err := checker.CheckGitVersion(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
	if err := checker.VerifyGoImportsRename(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
