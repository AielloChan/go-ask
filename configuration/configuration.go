package configuration

import (
	"errors"
	"os"
	"path/filepath"
)

func FindStageIndexByName(stages *[]Stage, name string) (int, bool) {
	if name == "" {
		return -1, false
	}
	for i, item := range *stages {
		if item.Name == name {
			return i, true
		}
	}
	return -1, false
}

func InitConfig(filePth string) (Configuration, error) {
	curPath, _ := os.Getwd()
	_, err := os.Stat(filepath.Join(curPath, filePth))
	if err != nil {
		return Configuration{}, errors.New("Config 'ask.config.json' not exist at current path '" + curPath + "' or do not have access permission")
	}

	cfg, err := GetConfig(filePth)
	if err != nil {
		return Configuration{}, errors.New("Error while read config: " + err.Error())
	}
	if len(cfg.Stages) == 0 {
		return cfg, errors.New("Does not have 'stages' config")
	}
	return cfg, nil
}
