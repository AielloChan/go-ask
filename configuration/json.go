package configuration

import (
	"encoding/json"
	"os"
)

type SelectOptions struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Next  string `json:"next"`
}

type StageConfig struct {
	Min     int             `json:"min"`
	Max     int             `json:"max"`
	Size    int             `json:"size"`
	Options []SelectOptions `json:"options"`
	Default interface{}     `json:"default"`
	Next    string          `json:"next"`
	Cmd     string          `json:"cmd"`
	Success string          `json:"success"`
	Failed  string          `json:"failed"`
}

type Stage struct {
	Label  string      `json:"label"`
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Config StageConfig `json:"config"`
	Next   string      `json:"next"`
}

type Configuration struct {
	Stages []Stage `json:"stages"`
}

func GetConfig(cfgPath string) (Configuration, error) {
	file, _ := os.Open(cfgPath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err := decoder.Decode(&conf)
	return conf, err
}
