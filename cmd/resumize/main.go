package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/PappasBrent/resumize/internal/args"
	"github.com/PappasBrent/resumize/internal/cv"
	"github.com/PappasBrent/resumize/internal/resumewriter"
)

// TODO: Add references
// TODO: Add explanation on how to use the program
func main() {
	args, err := args.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	cv, err := cv.ReadFile(args.CVDataYAMLFilepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	templ, err := template.ParseFiles(args.TemplateFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	rw := resumewriter.ResumeWriter{CV: cv, Template: templ}

	err = rw.WriteFiles(args.OutputFilePaths)
	if err != nil {
		fmt.Println(err)
		return
	}
}
