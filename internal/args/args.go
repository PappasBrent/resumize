package args

type Args struct {
	CVDataYAMLFilepath string
	TemplateFilePath   string
	OutputFilePaths    []string
}

// Parse parses a list of command line arguments
// (not including the path to the program)
func Parse(args []string) (*Args, error) {
	if n := len(args); n < 1 {
		return nil, &ArgsError{"No CV data file provided"}
	} else if n < 2 {
		return nil, &ArgsError{"No template file provided"}
	} else if n < 3 {
		return nil, &ArgsError{"No output filepaths provided"}
	}
	return &Args{args[0], args[1], args[2:]}, nil
}
