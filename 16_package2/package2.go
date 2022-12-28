package main

import (
	_ "database/sql"          // 사용하지 않을 패키지 import
	htemplate "html/template" // 별칭으로 import
	"text/template"
)

func main() {
	template.New("foo").Parse(`{{define "T"}}Hello`)
	htemplate.New("foo").Parse(`{{define "T"}}Hello`)
}
