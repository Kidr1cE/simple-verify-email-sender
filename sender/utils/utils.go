package utils

import (
	"bytes"
	"fmt"
	"text/template"
)

func FromFormat(from, email string) string {
	return fmt.Sprintf("%s <%s>", from, email)
}

func GetTemplate(text string) (*template.Template, error) {
	tmpl, err := template.New("HTML").Parse(text)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func GetContent(tmpl *template.Template, table []string) []byte {
	var buf bytes.Buffer
	tmpl.Execute(&buf, table)
	return buf.Bytes()
}
