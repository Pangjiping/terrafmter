package gen

import (
	"bytes"
	"embed"
	"github.com/Pangjiping/terrafmtter/format"
	"io"
	"text/template"
)

//go:embed templates/*
var tmpl embed.FS

func Execute(fmtter format.Formatter, w io.Writer) error {
	templ := template.New("main.tftmpl")
	tpl, err := templ.ParseFS(tmpl, "templates/*.tftmpl")
	if err != nil {
		return err
	}
	buf := &bytes.Buffer{}
	if err = tpl.ExecuteTemplate(buf, "main.tftmpl", &fmtter); err != nil { // todo: wrong
		return err
	}

	if err = formatt(buf, w); err != nil {
		return err
	}
	return nil
}

func formatt(in io.Reader, out io.Writer) error {
	return nil
}
