package forms

import (
	"html/template"
	"strings"
)

func HTML(t *template.Template, strct interface{}) (template.HTML, error) {

	var inputs []string

	for _, field := range fields(strct) {
		var sb strings.Builder
		err := t.Execute(&sb, field)
		if err != nil {
			return "", err
		}
		inputs = append(inputs, sb.String())
	}
	return template.HTML(strings.Join(inputs, "")), nil
}
