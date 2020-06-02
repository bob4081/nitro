package validate

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Hostname(v string) error {
	msg := "you must provide a valid domain, without a TLD and only lowercase"

	if strings.Contains(v, " ") {
		return errors.New(msg)
	}

	lower := strings.ToLower(v)
	if lower != v {
		return errors.New(msg)
	}

	return nil
}

// path will check is a fali
func Path(p string) error {
	f, err := os.Stat(p)
	if err != nil {

		return err
	}

	if f.IsDir() {
		return nil
	}

	return errors.New("you must provide a path to a valid directory")
}

// PHPVersion takes a string that represents a PHP version to install and returns and error if that PHP version
// is not yet supported.
func PHPVersion(v string) error {
	switch v {
	case "7.4":
		return nil
	case "7.3":
		return nil
	case "7.2":
		return nil
	}

	return errors.New(fmt.Sprintf("the PHP version %q is not valid", v))
}

func Memory(v string) error {
	if !strings.HasSuffix(v, "G") {
		return errors.New("memory must end with a G")
	}

	return nil
}

func DiskSize(v string) error {
	if !strings.HasSuffix(v, "G") {
		return errors.New("disk must end with a G")
	}

	return nil
}

func MachineName(v string) error {
	if v == "" {
		return errors.New("machine name cannot be empty")
	}
	if strings.Contains(v, " ") {
		return errors.New("machine name cannot contain spaces")
	}

	return nil
}
