package cv

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func Test(t *testing.T) {

	may1st1999 := time.Date(1999, time.May, 1, 0, 0, 0, 0, time.UTC)
	april10th2021 := time.Date(2021, time.April, 10, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		desc     string
		data     string
		expected *CV
	}{
		{
			desc: "Spongebob's CV",
			data: `name: Spongebob Squarepants
summary: Sea Sponge | Frycook | Karate Master
contact:
  email: sponge.bob@nickelodeon.com
  location: Bikini Bottom
  phone: 123-456-7890
links:
  twitter: "@spongebob"
acronyms:
  KK: Krusty Krab
  EOM: Employee of the Month
skills:
  - name: Cooking
    level: 7
    notes:
      - I have years of experience at the Krusty Krab
  - name: Karate
    level: 6
    notes:
      - I regularly practice with the land mammal Sandy Cheeks
employment-history:
  - company: The Krusty Krab
    address: Bikini Bottom
    position: Frycook
    currently-employed: true
    start-date: 1999-05-01
    end-date: 2021-04-10
    responsibilities:
      - Cook burgers
      - Protect the Secret Formula
education-history:
  - institute: Bikini Bottom Jellyfishing Academy
    major: Jellyfishing
    gpa: 4.0
    currently-attending: false
    start-date: 1999-05-01
    end-date: 2021-04-10
    degree: Grand Jellyfisher
awards:
  - EOM at the KK for 144 consecutive months`,
			expected: &CV{
				Name:    "Spongebob Squarepants",
				Summary: "Sea Sponge | Frycook | Karate Master",
				Contact: Contact{
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
				Skills: []Skill{
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
				EmploymentHistory: []Employment{
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
				EducationHistory: []Education{
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
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			file, err := os.CreateTemp(".", "*.yaml")
			if err != nil {
				t.Errorf("Error creating file %q", file.Name())
			}
			_, err = file.WriteString(tC.data)
			if err != nil {
				t.Errorf("Error creating file %q", file.Name())
			}
			file.Close()

			result, err := ReadFile(file.Name())
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("Test %q\nOutput: %#v\n Expected: %#v", tC.desc, result, tC.expected)
			}

			err = os.Remove(file.Name())
			if err != nil {
				t.Errorf("Error removing temp file %q", file.Name())
			}
		})
	}
}
