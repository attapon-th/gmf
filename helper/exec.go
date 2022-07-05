package helper

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	spaceRegexp = regexp.MustCompile("[\\s]+")
)

// Exec cli commands
func Exec(command string, args ...string) (output string, err error) {
	commands := spaceRegexp.Split(command, -1)
	command = commands[0]
	commandArgs := []string{}
	if len(commands) > 1 {
		commandArgs = commands[1:]
	}
	if len(args) > 0 {
		commandArgs = append(commandArgs, args...)
	}

	fullCommand, err := exec.LookPath(command)
	if err != nil {
		return "", fmt.Errorf("%s cannot be found", command)
	}

	cmd := exec.Command(fullCommand, commandArgs...)
	cmd.Env = os.Environ()

	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr

	// logger.Debug(fullCommand, " ", strings.Join(commandArgs, " "))

	out, err := cmd.Output()
	if err != nil {
		// log.Error().Err(err).Msgf("%s %s", fullCommand, strings.Join(commandArgs, " "))
		err = errors.New(stdErr.String())
		return
	}

	output = strings.Trim(string(out), "\n")
	return
}

// Exec cli commands
// func ExecOutStd(command string, args ...string) (io.ReadCloser, error) {
// 	commands := spaceRegexp.Split(command, -1)
// 	command = commands[0]
// 	commandArgs := []string{}
// 	if len(commands) > 1 {
// 		commandArgs = commands[1:]
// 	}
// 	if len(args) > 0 {
// 		commandArgs = append(commandArgs, args...)
// 	}

// 	fullCommand, err := exec.LookPath(command)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s cannot be found", command)
// 	}

// 	cmd := exec.Command(fullCommand, commandArgs...)
// 	cmd.Env = os.Environ()
// 	cmd.Stderr = cmd.Stdout
// 	stdOut, err := cmd.StdoutPipe()
// 	if err != nil {
// 		return stdOut, err
// 	}
// 	if err := cmd.Run(); err != nil {
// 		return stdOut, err
// 	}
// 	return stdOut, nil
// }

func ConcatError(err ...error) error {
	errStr := ""
	for _, err := range err {
		errStr += err.Error()
	}
	return errors.New(errStr)
}
