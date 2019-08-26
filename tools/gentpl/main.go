package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"text/template"
)

type context struct {
	platform string
}

func getTargetPath(src string) string {
	basename := src[:strings.LastIndex(src, ".")]
	platform := runtime.GOOS
	return fmt.Sprintf("%s_%s.go", basename, platform)
}

func renderTemplate(tpl *template.Template, dst string) {
	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_TRUNC, 0)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	ctx := context{
		runtime.GOOS,
	}
	err = tpl.Execute(out, ctx)
	if err != nil {
		panic(err)
	}
}

func main() {
	filename := os.Args[1]
	template := template.Must(template.ParseFiles(filename))
	renderTemplate(template, getTargetPath(filename))
}
