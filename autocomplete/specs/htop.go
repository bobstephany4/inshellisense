// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["htop"] = model.Subcommand{
		Name:        []string{"htop"},
		Description: `Improved top (interactive process viewer)`,
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Show help for htop`,
		}, {
			Name:        []string{"--no-color", "-C"},
			Description: `Use a monochrome color scheme`,
		}, {
			Name:        []string{"--delay", "-d"},
			Description: `Delay between updates, in tenths of sec`,
			Args: []model.Arg{{
				Name:        "delay",
				Suggestions: []model.Suggestion{{Name: []string{`10`}}, {Name: []string{`1`}}, {Name: []string{`60`}}},
			}},
		}, {
			Name:        []string{"--filter", "-F"},
			Description: `Filter commands`,
			Args: []model.Arg{{
				Name: "filter",
			}},
		}, {
			Name:        []string{"--highlight-changes", "-H"},
			Description: `Highlight new and old processes`,
			Args: []model.Arg{{
				Name:        "delay",
				Description: `Delay between updates of highlights, in tenths of sec`,
				Suggestions: []model.Suggestion{{Name: []string{`10`}}, {Name: []string{`1`}}, {Name: []string{`60`}}},
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"--no-mouse", "-M"},
			Description: `Disable the mouse`,
		}, {
			Name:        []string{"--pid", "-p"},
			Description: `Show only the given PIDs`,
			Args: []model.Arg{{
				Name:       "PID",
				IsVariadic: true,
			}},
		}, {
			Name:        []string{"--sort-key", "-s"},
			Description: `Sort by COLUMN in list view`,
			Args: []model.Arg{{
				Name: "column",
			}},
		}, {
			Name:        []string{"--tree", "-t"},
			Description: `Show the tree view`,
		}, {
			Name:        []string{"--user", "-u"},
			Description: `Show only processes for a given user (or $USER)`,
			Args: []model.Arg{{
				Name:        "user",
				Suggestions: []model.Suggestion{{Name: []string{`$USER`}}},
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"--no-unicode", "-U"},
			Description: `Do not use unicode but plain ASCII`,
		}, {
			Name:        []string{"--version", "-V"},
			Description: `Print version info`,
		}},
	}
}
