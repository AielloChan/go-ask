package tools

import (
	"bytes"
	"errors"
	"html/template"
	"strings"

	"aiellochan.com/go-ask/model"
)

func ProcessTpl(tplStr string, store *model.Store) (string, error) {
	var preview bytes.Buffer
	tpl, err := template.New("preview").Parse(tplStr)
	if err != nil {
		return "", errors.New("Parse template string '" + tplStr + "' error: " + err.Error())
	}
	err = tpl.Execute(&preview, &store)
	if err != nil {
		return "", errors.New("Execute template string '" + tplStr + "' error: " + err.Error())
	}
	return strings.ReplaceAll(preview.String(), "\\n", "\n"), nil
}
