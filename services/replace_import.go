package services

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strings"
)

func ReplaceImportFile(filename string, oldName, moduleName string) error {
	oldName = strings.TrimRight(strings.TrimSpace(oldName), "/") + "/"
	fs, err := os.OpenFile(filename, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer fs.Close()

	isImport := true
	count := 0
	bs := bufio.NewScanner(fs)
	for bs.Scan() {
		line := bs.Text()
		if strings.Contains(line, "import (") {
			isImport = true
			continue
		}
		if isImport {
			if strings.Contains(line, oldName) {
				count++
			} else if strings.Contains(strings.TrimSpace(line), ")") {
				break
			}
		}
	}
	if count == 0 {
		return nil
	}
	b, err := ioutil.ReadAll(fs)
	if err != nil {
		return err
	}

	moduleName = strings.TrimRight(moduleName, "/") + "/"
	b = bytes.Replace(b, []byte("gitlab.com/indev-moph/fiber-api/"), []byte(moduleName), count)
	_, err = fs.Write(b)
	return err
}
