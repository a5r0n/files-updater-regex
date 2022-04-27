package updater

import (
	"io/ioutil"
	"os"
	"regexp"
)

var FUVERSION = "dev"

type Updater struct {
	pattern *regexp.Regexp
}

func (u *Updater) Init(m map[string]string) error {
	if pattern, ok := m["pattern"]; ok {
		var err error
		u.pattern, err = regexp.Compile(pattern)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *Updater) Name() string {
	return "regex"
}

func (u *Updater) Version() string {
	return FUVERSION
}

func (u *Updater) ForFiles() string {
	return ".*"
}

func (u *Updater) writeFile(fName string, newVersion string) error {
	file, err := os.OpenFile(fName, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	if !u.pattern.Match(buf) {
		return os.ErrInvalid
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	newFileContent := u.pattern.ReplaceAllString(string(buf), newVersion)
	_, err = file.WriteString(newFileContent)
	return err
}

func (u *Updater) Apply(file, newVersion string) error {
	return u.writeFile(file, newVersion)
}
