package args

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc       string
		args       []string
		shouldFail bool
		expected   *Args
	}{
		{
			desc:       "No CV data file",
			args:       []string{},
			shouldFail: true,
			expected:   nil,
		},
		{
			desc:       "No template file",
			args:       []string{"CV.yaml"},
			shouldFail: true,
			expected:   nil,
		},
		{
			desc:       "No output filepaths",
			args:       []string{"CV.yaml", "template.templ"},
			shouldFail: true,
			expected:   nil,
		},
		{
			desc:       "Success",
			args:       []string{"CV.yaml", "template.templ", "output.html"},
			shouldFail: false,
			expected:   &Args{"CV.yaml", "template.templ", []string{"output.html"}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			args, err := Parse(tC.args)
			if tC.shouldFail {
				if err == nil {
					t.Errorf("Test %q did not fail when it was expected to", tC.desc)
				}
				if _, ok := err.(*ArgsError); !ok {
					t.Errorf("Test %q failed, but not due to an ArgsError as expected", tC.desc)
				}
				return
			}
			if !reflect.DeepEqual(args, tC.expected) {
				t.Errorf("Test %q\nOutput: %v Expected: %v", tC.desc, args, tC.expected)
			}
		})
	}
}
