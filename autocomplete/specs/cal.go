// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["cal"] = model.Subcommand{
		Name:        []string{"cal"},
		Description: `Displays a calendar and the date of Easter`,
		Args: []model.Arg{{
			Name:        "year",
			Description: `Year to print calendar of`,
		}},
		Options: []model.Option{{
			Name:        []string{"-h"},
			Description: `Turns off highlighting of today`,
		}, {
			Name:        []string{"-j"},
			Description: `Display Julian days (days one-based, numbered from January 1)`,
		}, {
			Name:        []string{"-m"},
			Description: `Display the specified month.  If month is specified as a decimal number, it may be followed by the letter ‘f’ or ‘p’ to indicate the following or preceding month of that number, respectively`,
			Args: []model.Arg{{
				Name: "month",
				Suggestions: []model.Suggestion{{
					Name: []string{`january`},
				}, {
					Name: []string{`february`},
				}, {
					Name: []string{`march`},
				}, {
					Name: []string{`april`},
				}, {
					Name: []string{`may`},
				}, {
					Name: []string{`june`},
				}, {
					Name: []string{`july`},
				}, {
					Name: []string{`august`},
				}, {
					Name: []string{`september`},
				}, {
					Name: []string{`october`},
				}, {
					Name: []string{`november`},
				}, {
					Name: []string{`december`},
				}},
			}},
			ExclusiveOn: []string{"-y"},
		}, {
			Name:        []string{"-y"},
			Description: `Display a calendar for the specified year`,
			ExclusiveOn: []string{"-m"},
		}},
	}
}
