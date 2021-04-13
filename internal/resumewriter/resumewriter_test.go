package resumewriter

import (
	"os"
	"testing"
	"text/template"
	"time"

	"github.com/PappasBrent/resumize/internal/cv"
)

func Test(t *testing.T) {
	may1st1999 := time.Date(1999, time.May, 1, 0, 0, 0, 0, time.UTC)
	april10th2021 := time.Date(2021, time.April, 10, 0, 0, 0, 0, time.UTC)

	spongebobCV := &cv.CV{
		Name:    "Spongebob Squarepants",
		Summary: "Sea Sponge | Frycook | Karate Master",
		Contact: cv.Contact{
			Email:    "sponge.bob@nickelodeon.com",
			Location: "Bikini Bottom",
			Phone:    "123-456-7890",
		},
		Links: map[string]string{
			"twitter": "@spongebob",
		},
		Acronyms: map[string]string{
			"KK":  "Krusty Krab",
			"EOM": "Employee of the Month",
		},
		Skills: []cv.Skill{
			{
				Name:  "Cooking",
				Level: 7,
				Notes: []string{"I have years of experience at the Krusty Krab"},
			},
			{
				Name:  "Karate",
				Level: 6,
				Notes: []string{"I regularly practice with the land mammal Sandy Cheeks"},
			},
		},
		EmploymentHistory: []cv.Employment{
			{
				Company:           "The Krusty Krab",
				Address:           "Bikini Bottom",
				Position:          "Frycook",
				CurrentlyEmployed: true,
				StartDate:         may1st1999,
				EndDate:           april10th2021,
				Responsibilities:  []string{"Cook burgers", "Protect the Secret Formula"},
			},
		},
		EducationHistory: []cv.Education{
			{
				Institute:          "Bikini Bottom Jellyfishing Academy",
				Major:              "Jellyfishing",
				GPA:                4.0,
				CurrentlyAttending: false,
				StartDate:          may1st1999,
				EndDate:            april10th2021,
				Degree:             "Grand Jellyfisher",
			},
		},
		Awards: []string{
			"EOM at the KK for 144 consecutive months",
		},
	}

	testCases := []struct {
		desc      string
		rw        ResumeWriter
		filepaths []string
		expected  string
	}{
		{
			desc: "Name, 1 file",
			rw: ResumeWriter{
				CV: spongebobCV,
				Template: template.Must(template.New("test").Parse(
					`<h1>{{.Name}}</h1>`,
				)),
			},
			filepaths: []string{"test.html"},
			expected:  "<h1>Spongebob Squarepants</h1>",
		},
		{
			desc: "Name, 2 files",
			rw: ResumeWriter{
				CV: spongebobCV,
				Template: template.Must(template.New("test").Parse(
					`<h1>{{.Name}}</h1>`,
				)),
			},
			filepaths: []string{"test-1.html", "test-2.html"},
			expected:  "<h1>Spongebob Squarepants</h1>",
		},
		{
			desc: "Full Resume",
			rw: ResumeWriter{
				CV: spongebobCV,
				Template: template.Must(template.New("spongebob-resume-test").Parse(
					`<header>{{.Name}}</header>
<small>{{.Summary}}</small>
<section>
    <header>Contact</header>
    <ul>
        <li>Email: {{.Contact.Email}}</li>
        <li>Location: {{.Contact.Location}}</li>
        <li>Phone: {{.Contact.Phone}}</li>
    </ul>
</section>
<section>
    <header>Links</header>
    <ul>
        <li>Twitter: {{.Links.twitter}}</li>
    </ul>
</section>
<section>
    <header>Acronyms</header>
    <ul>{{range $letters, $phrase := .Acronyms}}
        <li>{{$letters}}: {{$phrase}}</li>{{end}}
    </ul>
</section>
<section>
    <header>Skills</header>{{range $skill := .Skills}}
    <div>
        <h2>{{$skill.Name}}</h2>
        <p>Skill level: {{$skill.Level}}</p>
        <ul>{{range $note := $skill.Notes}}
            <li>{{$note}}</li>{{end}}
        </ul>
    </div>{{end}}
</section>
<section>
    <header>Employment History</header>{{range $employment := .EmploymentHistory}}
    <div>
        <h2>Position: {{$employment.Position}}</h2>
        <p>Company: {{$employment.Company}}</p>
        <p>Address: {{$employment.Address}}</p>
        <p>Employed from: {{$employment.EmploymentDuration}}</p>
        <h3>Responsibilities</h3>
        <ul>{{range $responsibility := $employment.Responsibilities}}
            <li>{{$responsibility}}</li>{{end}}
        </ul>
    </div>{{end}}
</section>
<section>
    <header>Education History</header>{{range $education := .EducationHistory}}
    <div>
        <h2>Institute: {{$education.Institute}}</h2>
        <p>Major: {{$education.Major}}</p>
        <p>Degree: {{$education.Degree}}</p>
        <p>GPA: {{printf "%.2f" $education.GPA}}</p>
        <p>Attended from: {{$education.EducationDuration}}</p>
    </div>{{end}}
</section>
<section>
    <Header>Awards</Header>
    <ul>{{range $award := .Awards}}
        <li>{{$award}}</li>{{end}}
    </ul>
</section>`,
				)),
			},
			filepaths: []string{"resume.html"},
			expected: `<header>Spongebob Squarepants</header>
<small>Sea Sponge | Frycook | Karate Master</small>
<section>
    <header>Contact</header>
    <ul>
        <li>Email: sponge.bob@nickelodeon.com</li>
        <li>Location: Bikini Bottom</li>
        <li>Phone: 123-456-7890</li>
    </ul>
</section>
<section>
    <header>Links</header>
    <ul>
        <li>Twitter: @spongebob</li>
    </ul>
</section>
<section>
    <header>Acronyms</header>
    <ul>
        <li>EOM: Employee of the Month</li>
        <li>KK: Krusty Krab</li>
    </ul>
</section>
<section>
    <header>Skills</header>
    <div>
        <h2>Cooking</h2>
        <p>Skill level: 7</p>
        <ul>
            <li>I have years of experience at the Krusty Krab</li>
        </ul>
    </div>
    <div>
        <h2>Karate</h2>
        <p>Skill level: 6</p>
        <ul>
            <li>I regularly practice with the land mammal Sandy Cheeks</li>
        </ul>
    </div>
</section>
<section>
    <header>Employment History</header>
    <div>
        <h2>Position: Frycook</h2>
        <p>Company: The Krusty Krab</p>
        <p>Address: Bikini Bottom</p>
        <p>Employed from: May 1, 1999 - Present</p>
        <h3>Responsibilities</h3>
        <ul>
            <li>Cook burgers</li>
            <li>Protect the Secret Formula</li>
        </ul>
    </div>
</section>
<section>
    <header>Education History</header>
    <div>
        <h2>Institute: Bikini Bottom Jellyfishing Academy</h2>
        <p>Major: Jellyfishing</p>
        <p>Degree: Grand Jellyfisher</p>
        <p>GPA: 4.00</p>
        <p>Attended from: May 1, 1999 - April 10, 2021</p>
    </div>
</section>
<section>
    <Header>Awards</Header>
    <ul>
        <li>EOM at the KK for 144 consecutive months</li>
    </ul>
</section>`,
		},
		{
			// TODO: Check this test, finish adding markdown functionality
			desc: "Markup",
			rw: ResumeWriter{
				CV: &cv.CV{
					Name: "[My website](https://pappasbrent.com)",
				},
				Template: template.Must(template.New("markup-test").Parse(
					`{{.Name}}`,
				))},
			filepaths: []string{"markup-test.html"},
			expected:  `<a href="https://pappasbrent.com">My website</a>`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := tC.rw.WriteFiles(tC.filepaths)
			if err != nil {
				t.Errorf("Test %q: %v", tC.desc, err)
			}
			for _, fp := range tC.filepaths {
				data, err := os.ReadFile(fp)
				if err != nil {
					t.Errorf("Error opening file %q: %v", fp, err)
				}
				if output := string(data); output != tC.expected {
					t.Errorf("Test %q\nOutput: %v\nExpected: %v", tC.desc, output, tC.expected)
				}
				// This remove will fail if the test fails,
				// technically a bug but useful for reviewing failed output
				if err = os.Remove(fp); err != nil {
					t.Errorf("Error removing file %q: %v", fp, err)
				}
			}
		})
	}
}
