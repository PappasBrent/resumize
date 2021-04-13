package resumewriter

import (
	"os"
	"text/template"

	"github.com/PappasBrent/resumize/internal/cv"
)

type ResumeWriter struct {
	CV       *cv.CV
	Template *template.Template
}

// WriteFiles writes the formatted resume data to each file in filepaths
func (rw *ResumeWriter) WriteFiles(filepaths []string) error {
	cvhlv := cv.CVHyperLinkVisitor{}
	rw.CV.Accept(&cvhlv)
	for _, fp := range filepaths {
		file, err := os.Create(fp)
		if err != nil {
			return err
		}
		if err = rw.Template.Execute(file, rw.CV); err != nil {
			return err
		}
		if err = file.Close(); err != nil {
			return err
		}
	}
	return nil
}
