// static.v2
// Copyright (c) 2020, HuguesGuilleus
// BSD 3-Clause License

package static

import (
	"html/template"
	// "io/ioutil"
	// "log"
	// "os"
	// "path/filepath"
)

// Create template from b of the file f.
func TemplateHTML(b []byte, f string) *template.Template {
	t := template.New("")

	var min Minifier = HtmlMinify

	go func() {
		t.Parse(min.minS(b))
		for range ticker() {
			if content := readFileOnce(f, min); content != nil {
				t.Parse(string(content))
			}
		}
	}()

	return t
}
