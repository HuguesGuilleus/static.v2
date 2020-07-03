// static.v2
// Copyright (c) 2020, HuguesGuilleus
// BSD 3-Clause License

package static

import (
	templatehtml "html/template"
	templatetext "text/template"
)

// Create template from b or the file f.
func Template(b []byte, f string, min Minifier) *templatetext.Template {
	t := templatetext.New("")

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

// Create template from b or the file f.
func TemplateHTML(b []byte, f string) *templatehtml.Template {
	t := templatehtml.New("")

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
