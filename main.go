package main

import (
	"flag"
	"fmt"
	"os"

	"aiellochan.com/go-ask/configuration"
	"aiellochan.com/go-ask/model"
	"aiellochan.com/go-ask/pipline"
	"aiellochan.com/go-ask/tools"
	"github.com/AlecAivazis/survey/v2/terminal"
)

var (
	eject  bool
	config string
)

func init() {
	flag.BoolVar(&eject, "eject", false, "eject default config file")
	flag.StringVar(&config, "config", "ask.config.json", "string flag value")
}

func run(configFile string) {
	// Create data store
	store := model.GetInstance()

	// Get config
	cfg, err := configuration.InitConfig(configFile)
	if err != nil {
		tools.PrintError(err.Error() + "\n")
		os.Exit(0)
	}

	// Run pipline
	err = pipline.RunJob(&cfg.Stages, 0, store)

	// Ctrl + c
	if err == terminal.InterruptErr {
		os.Exit(0)
	} else if err != nil {
		tools.PrintError(err.Error() + "\n")
		os.Exit(0)
	}
}

func main() {
	flag.Parse()
	if eject {
		// eject config
		fmt.Println("Comming soon...")
	} else {
		run(config)
	}
}
