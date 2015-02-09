package ieversionlocker

import (
	"fmt"
	"strings"
	"syscall"

	"github.com/hnakamur/w32registry"
)

type RegValueName struct {
	key     syscall.Handle
	subkey  string
	valname string
}

type IEVersion int

const (
	IE7 = iota + 7
	IE8
	IE9
	IE10
	IE11
)

var regValueNames = map[IEVersion]RegValueName{
	IE8: RegValueName{
		key:     syscall.HKEY_LOCAL_MACHINE,
		subkey:  `SOFTWARE\Microsoft\Internet Explorer\Setup\8.0`,
		valname: "DoNotAllowIE80",
	},
	IE9: RegValueName{
		key:     syscall.HKEY_LOCAL_MACHINE,
		subkey:  `SOFTWARE\Microsoft\Internet Explorer\Setup\9.0`,
		valname: "DoNotAllowIE90",
	},
	IE10: RegValueName{
		key:     syscall.HKEY_LOCAL_MACHINE,
		subkey:  `SOFTWARE\Microsoft\Internet Explorer\Setup\10.0`,
		valname: "DoNotAllowIE10",
	},
	IE11: RegValueName{
		key:     syscall.HKEY_LOCAL_MACHINE,
		subkey:  `SOFTWARE\Microsoft\Internet Explorer\Setup\11.0`,
		valname: "DoNotAllowIE11",
	},
}

var newerVersions = map[IEVersion][]IEVersion{
	IE7:  []IEVersion{IE8, IE9, IE10, IE11},
	IE8:  []IEVersion{IE9, IE10, IE11},
	IE9:  []IEVersion{IE10, IE11},
	IE10: []IEVersion{IE11},
	IE11: []IEVersion{},
}

func CurrentVersion() (IEVersion, error) {
	fullVersion, err := w32registry.GetValueString(
		syscall.HKEY_LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Internet Explorer`,
		"svcVersion")
	if err == syscall.ERROR_FILE_NOT_FOUND {
		fullVersion, err = w32registry.GetValueString(
			syscall.HKEY_LOCAL_MACHINE,
			`SOFTWARE\Microsoft\Internet Explorer`,
			"Version")
	}
	if err != nil {
		return 0, err
	}

	i := strings.IndexRune(fullVersion, '.')
	if i == -1 {
		return 0, fmt.Errorf("Invalid IE version format: %s", fullVersion)
	}
	majorVersion := fullVersion[:i]
	switch majorVersion {
	case "7":
		return IE7, nil
	case "8":
		return IE8, nil
	case "9":
		return IE9, nil
	case "10":
		return IE10, nil
	case "11":
		return IE11, nil
	default:
		return 0, fmt.Errorf("Unknown IE major version: %s", fullVersion)
	}
}

func Lock(version IEVersion) error {
	var err error
	for _, v := range newerVersions[version] {
		err = Block(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func Unlock(version IEVersion) error {
	var err error
	for _, v := range newerVersions[version] {
		err = Unblock(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func Block(version IEVersion) error {
	n := regValueNames[version]
	return w32registry.SetKeyValueUint32(n.key, n.subkey, n.valname, 1)
}

func Unblock(version IEVersion) error {
	n := regValueNames[version]
	err := w32registry.DeleteKeyValue(n.key, n.subkey, n.valname)
	if err == syscall.ERROR_FILE_NOT_FOUND {
		err = nil
	}
	return err
}
