package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"text/template"
)

var (
	help bool
	cef  string
)

func init() {
	flag.BoolVar(&help, "help", false, "print help message")
	flag.StringVar(&cef, "cef", "", "CEF prebuilt dirname")
}

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
		CEFBuild: cef,
	}
	err = tpl.Execute(out, ctx)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// handle errors
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	// parse flags
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	if flag.NArg() < 1 {
		panic(errors.New("Not enough arguments"))
	}
	filename := flag.Arg(0)
	template := template.Must(template.ParseFiles(filename))
	err := renderTemplate(template, getTargetPath(filename))
	if err != nil {
		panic(err)
	}
}
