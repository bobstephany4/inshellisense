// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["uniq"] = model.Subcommand{
		Name:        []string{"uniq"},
		Description: `Report or omit repeated line`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFilepaths, model.TemplateFolders},
			Name:       "input",
			IsOptional: true,
		}, {
			Name:       "output",
			IsOptional: true,
		}},
		Options: []model.Option{{
			Name:        []string{"-c", "--count"},
			Description: `Prefix lines by the number of occurrences`,
		}, {
			Name:        []string{"-d", "--repeated"},
			Description: `Only print duplicate lines`,
		}, {
			Name:        []string{"-D", "--all-repeated"},
			Description: `Print all duplicate lines. Delimiting is done with blank lines`,
			Args: []model.Arg{{
				Name:        "delimit-method",
				Suggestions: []model.Suggestion{{Name: []string{`none`}}, {Name: []string{`prepend`}}, {Name: []string{`separate`}}},
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"-f", "--skip-fields"},
			Description: `Avoid comparing the first N fields`,
			Args: []model.Arg{{
				Name: "number",
			}},
		}, {
			Name:        []string{"-i", "--ignore-case"},
			Description: `Ignore differences in case when comparing`,
		}, {
			Name:        []string{"-s", "--skip-chars"},
			Description: `Avoid comparing the first N characters`,
			Args: []model.Arg{{
				Name: "number",
			}},
		}, {
			Name:        []string{"-u", "--unique"},
			Description: `Only print unique lines`,
		}, {
			Name:        []string{"-z", "--zero-terminated"},
			Description: `End lines with 0 byte, not newline`,
		}, {
			Name:        []string{"-w", "--check-chars"},
			Description: `Compare no more than N characters in lines`,
			Args: []model.Arg{{
				Name: "number",
			}},
		}, {
			Name:        []string{"--help"},
			Description: `Display this help and exit`,
		}, {
			Name:        []string{"--version"},
			Description: `Output version information and exit`,
		}},
	}
}
