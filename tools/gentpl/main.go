package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"text/template"
)

type context struct {
	Platform string
	CEFBuild string
}

func getTargetPath(src string) string {
	basename := src[:strings.LastIndex(src, ".")]
	platform := runtime.GOOS
	return fmt.Sprintf("%s_%s.go", basename, platform)
}

func renderTemplate(tpl *template.Template, dst string) error {
	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_TRUNC, 0)
	if err != nil {
		return err
	}
	defer out.Close()
	ctx := context{
		Platform: runtime.GOOS,
		CEFBuild: "cef",
	}
	err = tpl.Execute(out, ctx)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	filename := os.Args[1]
	template := template.Must(template.ParseFiles(filename))
	err := renderTemplate(template, getTargetPath(filename))
	if err != nil {
		panic(err)
	}
}
