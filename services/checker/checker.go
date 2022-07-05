package checker

import (
	"errors"
	"fmt"
	"strings"

	"github.com/attapon-th/gmf/helper"
	"github.com/phuslu/log"
)

const (
	REQ_GO_VERSION = "1.8.0"
)

func CheckGoVersion() error {
	output, err := helper.Exec("go", "version")
	if err != nil {
		return err
	}
	fmt.Println(output)
	if !strings.Contains(output, "go1.18.") {
		return errors.New("go version 1.18.x or more, https://go.dev/dl/")
	}
	return nil
}

func CheckGitVersion() error {
	output, err := helper.Exec("git", "version")
	if err != nil {
		return err
	}
	fmt.Println(output)
	if !strings.Contains(output, "git version ") {
		return errors.New("pleace install git. https://git-scm.com/downloads")
	}
	return nil
}

func VerifyGoImportsRename() error {
	_, err := helper.Exec("go-imports-rename")
	if !strings.Contains(err.Error(), "Usage: go-imports-rename") {
		log.Error().Err(err).Msg("")
		return installGoImportsRename()
	}
	fmt.Println(err.Error())
	return nil
}

func installGoImportsRename() error {
	_, err := helper.Exec("go", "get", "github.com/sirkon/go-imports-rename")
	if err != nil {
		return fmt.Errorf(err.Error(), "\nCan't `go get github.com/sirkon/go-imports-rename`, please recheck manual.")
	}
	fmt.Println("go", "get", "github.com/sirkon/go-imports-rename")

	_, err = helper.Exec("go", "install", "github.com/sirkon/go-imports-rename")
	if err != nil {
		return fmt.Errorf(err.Error(), "Can't `go install github.com/sirkon/go-imports-rename`, please recheck manual.")
	}
	fmt.Println("go", "install", "github.com/sirkon/go-imports-rename")
	return nil
}
