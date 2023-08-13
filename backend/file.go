package backend

import (
	"os"
	"path"

	"github.com/adrg/xdg"
)

func LocateConfigFile(configFilename string, portable bool) (string, error) {
	if portable {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return path.Join(cwd, configFilename), nil
	}
	return xdg.ConfigFile(configFilename)
}

func LocateDataFile(configFilename string, portable bool) (string, error) {
	if portable {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return path.Join(cwd, configFilename), nil
	}
	return xdg.DataFile(configFilename)
}
