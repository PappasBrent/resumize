package cv

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"gopkg.in/yaml.v2"
)

type CV struct {
	Name              string
	Summary           string
	Contact           Contact
	Links             map[string]string
	Acronyms          map[string]string
	Skills            []Skill
	EmploymentHistory []Employment `yaml:"employment-history"`
	EducationHistory  []Education  `yaml:"education-history"`
	Awards            []string
}

type Contact struct {
	Email    string
	Location string
	Phone    string
}

type Skill struct {
	Name  string
	Level int
	Notes []string
}

type Employment struct {
	Company           string
	Address           string
	Position          string
	CurrentlyEmployed bool      `yaml:"currently-employed"`
	StartDate         time.Time `yaml:"start-date"`
	EndDate           time.Time `yaml:"end-date"`
	Responsibilities  []string
}

type Education struct {
	Institute          string
	Major              string
	GPA                float64
	CurrentlyAttending bool
	StartDate          time.Time `yaml:"start-date"`
	EndDate            time.Time `yaml:"end-date"`
	Degree             string
}

// CVVisitor is an interface for CV visitors
type CVVisitor interface {
	VisitCV(*CV)
	VisitContact(*Contact)
	VisitSkills([]Skill)
	VisitEmploymentHistory([]Employment)
	VisitEducationHistory([]Education)
}

func (cv *CV) Accept(cvv CVVisitor) {
	cvv.VisitCV(cv)
	cvv.VisitContact(&cv.Contact)
	cvv.VisitSkills(cv.Skills)
	cvv.VisitEmploymentHistory(cv.EmploymentHistory)
	cvv.VisitEducationHistory(cv.EducationHistory)
}

// TODO: Use an embedded type with an interface
// for the Education and Employment?

// EmploymentDuration returns the formatted date range of employment
// TODO: Unit test
func (e *Employment) EmploymentDuration() string {
	dateFormat := "January 2, 2006"
	startString := e.StartDate.Format(dateFormat)
	if e.CurrentlyEmployed {
		return fmt.Sprintf("%s - Present", startString)
	}
	endString := e.EndDate.Format(dateFormat)
	return fmt.Sprintf("%s - %s", startString, endString)
}

// EducationDuration returns the formatted date range of education
// TODO: Unit test
func (e *Education) EducationDuration() string {
	dateFormat := "January 2, 2006"
	startString := e.StartDate.Format(dateFormat)
	if e.CurrentlyAttending {
		return fmt.Sprintf("%s - Present", startString)
	}
	endString := e.EndDate.Format(dateFormat)
	return fmt.Sprintf("%s - %s", startString, endString)
}

type CVHyperLinkVisitor struct{}

func replaceLinksWithAnchors(text string) string {
	pattern := regexp.MustCompile(`\[(?P<innerText>.*)\]\((?P<href>.*)\)`)
	anchorTemplate := `<a href="$href">$innerText</a>`
	return pattern.ReplaceAllStringFunc(text, func(s string) string {
		result := []byte{}
		for _, submatches := range pattern.FindAllStringSubmatchIndex(s, -1) {
			result = pattern.ExpandString(result, anchorTemplate, s, submatches)
		}
		return string(result)
	})
}

func (cvhlv *CVHyperLinkVisitor) VisitCV(cv *CV) {
	cv.Name = replaceLinksWithAnchors(cv.Name)
	cv.Summary = replaceLinksWithAnchors(cv.Summary)
	cvhlv.VisitContact(&cv.Contact)
	for linkName, link := range cv.Links {
		cv.Links[linkName] = replaceLinksWithAnchors(link)
	}
	for letters, acronym := range cv.Acronyms {
		cv.Acronyms[letters] = replaceLinksWithAnchors(acronym)
	}
	cvhlv.VisitSkills(cv.Skills)
	cvhlv.VisitEmploymentHistory(cv.EmploymentHistory)
	cvhlv.VisitEducationHistory(cv.EducationHistory)
	for i, award := range cv.Awards {
		cv.Awards[i] = replaceLinksWithAnchors(award)
	}
}

func (cvhlv *CVHyperLinkVisitor) VisitContact(contact *Contact) {
	contact.Email = replaceLinksWithAnchors(contact.Email)
	contact.Location = replaceLinksWithAnchors(contact.Location)
	contact.Phone = replaceLinksWithAnchors(contact.Phone)
}

func (cvhlv *CVHyperLinkVisitor) VisitSkills(skills []Skill) {
	for i, skill := range skills {
		skill.Name = replaceLinksWithAnchors(skill.Name)
		for j, note := range skill.Notes {
			skills[i].Notes[j] = replaceLinksWithAnchors(note)
		}
	}
}

func (cvhlv *CVHyperLinkVisitor) VisitEmploymentHistory(employmentHistory []Employment) {
	for i, employment := range employmentHistory {
		employmentHistory[i].Address = replaceLinksWithAnchors(employment.Address)
		employmentHistory[i].Company = replaceLinksWithAnchors(employment.Company)
		employmentHistory[i].Position = replaceLinksWithAnchors(employment.Position)
		for j, responsibility := range employment.Responsibilities {
			employmentHistory[i].Responsibilities[j] = replaceLinksWithAnchors(responsibility)
		}
	}
}

func (cvhlv *CVHyperLinkVisitor) VisitEducationHistory(educationHistory []Education) {
	for i, education := range educationHistory {
		educationHistory[i].Degree = replaceLinksWithAnchors(education.Degree)
		educationHistory[i].Institute = replaceLinksWithAnchors(education.Institute)
	}
}

// Read reads a CV data YAML string into a CV struct
func ReadFile(fp string) (*CV, error) {
	result := &CV{}
	data, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	yaml.Unmarshal(data, result)
	return result, nil
}
