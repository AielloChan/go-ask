package tools

import (
	"strings"

	"aiellochan.com/go-ask/model"
	"aiellochan.com/go-ask/shell"
)

const (
	ShellPrefix             = "#!"
	TemplatePrefix          = "#$"
	ShellThenTemplatePrefix = "#!$"
	TemplateThenShellPrefix = "#$!"
)

func ExecCommand(s string, store *model.Store) (string, error) {
	switch {
	case strings.HasPrefix(s, ShellThenTemplatePrefix):
		mixedStr := s[len(ShellThenTemplatePrefix):]
		shellRes, err := shell.ProcessShell(mixedStr)
		if err != nil {
			return "", err
		}
		res, err := ProcessTpl(shellRes, store)
		if err != nil {
			return "", err
		}
		return res, nil
	case strings.HasPrefix(s, TemplateThenShellPrefix):
		mixedStr := s[len(TemplateThenShellPrefix):]
		tplRes, err := ProcessTpl(mixedStr, store)
		if err != nil {
			return "", err
		}
		res, err := shell.ProcessShell(tplRes)
		if err != nil {
			return "", err
		}
		return res, nil
	case strings.HasPrefix(s, ShellPrefix):
		cmd := s[len(ShellPrefix):]
		res, err := shell.ProcessShell(cmd)
		if err != nil {
			return "", err
		}
		return res, nil
	case strings.HasPrefix(s, TemplatePrefix):
		tpl := s[len(TemplatePrefix):]
		return ProcessTpl(tpl, store)
	default:
		return s, nil
	}
}
