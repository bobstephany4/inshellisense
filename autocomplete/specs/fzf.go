// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["fzf"] = model.Subcommand{
		Name:        []string{"fzf"},
		Description: `A general-purpose command-line fuzzy finder`,
		Options: []model.Option{{
			Name:        []string{"-x", "--extended"},
			Description: `Enables extended-search mode`,
			ExclusiveOn: []string{"+x", "--no-extended"},
		}, {
			Name:        []string{"+x", "--no-extended"},
			Description: `Disables extended-search mode`,
			ExclusiveOn: []string{"-x", "--extended"},
		}, {
			Name:        []string{"-e", "--exact"},
			Description: `Enables Exact-match`,
		}, {
			Name:        []string{"--algo"},
			Description: `Fuzzy matching algorithm`,
			Args: []model.Arg{{
				Name:        "type",
				Suggestions: []model.Suggestion{{Name: []string{`v1`}}, {Name: []string{`v2`}}},
			}},
		}, {
			Name:        []string{"-i"},
			Description: `Case-insensitive match (default: smart-case match)`,
			ExclusiveOn: []string{"+i"},
		}, {
			Name:        []string{"+i"},
			Description: `Case-sensitive match (default: smart-case match)`,
			ExclusiveOn: []string{"-i"},
		}, {
			Name:        []string{"--literal"},
			Description: `Do not normalize latin script letters before matching`,
		}, {
			Name:        []string{"-n", "--nth"},
			Description: `Comma-separated list of field index expressions for limiting search scope`,
			Args: []model.Arg{{
				Name:        "index expressions",
				Description: `Non-zero integer or range expression ([BEGIN]..[END])`,
			}},
		}, {
			Name:        []string{"--with-nth"},
			Description: `Transform the presentation of each line using field index expressions`,
			Args: []model.Arg{{
				Name:        "index expressions",
				Description: `Non-zero integer or range expression ([BEGIN]..[END])`,
			}},
		}, {
			Name:        []string{"-d", "--delimiter"},
			Description: `Field delimiter regex (default: AWK-style)`,
			Args: []model.Arg{{
				Name: "STR",
			}},
		}, {
			Name:        []string{"+s", "--no-sort"},
			Description: `Do not sort the result`,
		}, {
			Name:        []string{"--tac"},
			Description: `Reverse the order of the input`,
		}, {
			Name:        []string{"--disabled"},
			Description: `Do not perform search`,
		}, {
			Name:        []string{"--tiebreak"},
			Description: `Comma-separated list of sort criteria to apply when the scores are tied`,
			Args: []model.Arg{{
				Name:        "criteria",
				Suggestions: []model.Suggestion{{Name: []string{`length`}}, {Name: []string{`begin`}}, {Name: []string{`end`}}, {Name: []string{`index`}}},
			}},
		}, {
			Name:        []string{"-m", "--multi"},
			Description: `Enable multi-select with tab/shift-tab`,
			Args: []model.Arg{{
				Name:       "MAX",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--no-mouse"},
			Description: `Disable mouse`,
		}, {
			Name:        []string{"--bind"},
			Description: `Custom key bindings. Refer to the man page`,
			Args: []model.Arg{{
				Name: "keybinds",
			}},
		}, {
			Name:        []string{"--cycle"},
			Description: `Enable cyclic scroll`,
		}, {
			Name:        []string{"--keep-right"},
			Description: `Keep the right end of the line visible on overflow`,
		}, {
			Name:        []string{"--no-hscroll"},
			Description: `Disable horizontal scroll`,
		}, {
			Name:        []string{"--hscroll-off"},
			Description: `Number of screen columns to keep to the right of the highlighted substring`,
			Args: []model.Arg{{
				Name: "columns",
			}},
		}, {
			Name:        []string{"--filepath-word"},
			Description: `Make word-wise movements respect path separators`,
		}, {
			Name:        []string{"--jump-labels"},
			Description: `Label characters for jump and jump-accept`,
			Args: []model.Arg{{
				Name: "characters",
			}},
		}, {
			Name:        []string{"--height"},
			Description: `Display fzf window below the cursor with the given height instead of using fullscreen`,
			Args: []model.Arg{{
				Name:        "height",
				Description: `Height[%]`,
			}},
		}, {
			Name:        []string{"--min-height"},
			Description: `Minimum height when --height is given in percent`,
			Args: []model.Arg{{
				Name: "height",
			}},
		}, {
			Name:        []string{"--layout"},
			Description: `Choose layout`,
			Args: []model.Arg{{
				Name:        "layout",
				Suggestions: []model.Suggestion{{Name: []string{`default`}}, {Name: []string{`reverse`}}, {Name: []string{`reverse-list`}}},
			}},
		}, {
			Name:        []string{"--border"},
			Description: `Draw border around the finder`,
			Args: []model.Arg{{
				Name:        "style",
				Suggestions: []model.Suggestion{{Name: []string{`rounded`}}, {Name: []string{`sharp`}}, {Name: []string{`horizontal`}}, {Name: []string{`vertical`}}, {Name: []string{`top`}}, {Name: []string{`bottom`}}, {Name: []string{`left`}}, {Name: []string{`right`}}, {Name: []string{`none`}}},
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"--margin"},
			Description: `Screen margin (TRBL | TB,RL | T,RL,B | T,R,B,L)`,
			Args: []model.Arg{{
				Name:        "margin",
				Description: `Number`,
			}},
		}, {
			Name:        []string{"--padding"},
			Description: `Padding inside border (TRBL | TB,RL | T,RL,B | T,R,B,L)`,
			Args: []model.Arg{{
				Name:        "padding",
				Description: `Number`,
			}},
		}, {
			Name:        []string{"--info"},
			Description: `Finder info style`,
			Args: []model.Arg{{
				Name:        "style",
				Suggestions: []model.Suggestion{{Name: []string{`default`}}, {Name: []string{`inline`}}, {Name: []string{`hidden`}}},
			}},
		}, {
			Name:        []string{"--prompt"},
			Description: `Input prompt`,
			Args: []model.Arg{{
				Name: "string",
			}},
		}, {
			Name:        []string{"--pointer"},
			Description: `Pointer to the current line`,
			Args: []model.Arg{{
				Name: "string",
			}},
		}, {
			Name:        []string{"--marker"},
			Description: `Multi-select marker`,
			Args: []model.Arg{{
				Name: "string",
			}},
		}, {
			Name:        []string{"--header"},
			Description: `String to print as header`,
			Args: []model.Arg{{
				Name: "string",
			}},
		}, {
			Name:        []string{"--header-lines"},
			Description: `The first N lines of the input are treated as header`,
			Args: []model.Arg{{
				Name: "number",
			}},
		}, {
			Name:        []string{"--ansi"},
			Description: `Enable processing of ANSI color codes`,
		}, {
			Name:        []string{"--tabstop"},
			Description: `Number of spaces for a tab character`,
			Args: []model.Arg{{
				Name: "spaces",
			}},
		}, {
			Name:        []string{"--color"},
			Description: `Base scheme`,
			Args: []model.Arg{{
				Name:        "color scheme",
				Description: `(dark|light|16|bw) and/or custom colors`,
				Suggestions: []model.Suggestion{{Name: []string{`dark`}}, {Name: []string{`light`}}, {Name: []string{`16`}}, {Name: []string{`bw`}}},
			}},
		}, {
			Name:        []string{"--no-bold"},
			Description: `Do not use bold text`,
		}, {
			Name:        []string{"--history"},
			Description: `History file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "file",
			}},
		}, {
			Name:        []string{"--history-size"},
			Description: `Maximum number of history entries`,
			Args: []model.Arg{{
				Name: "number",
			}},
		}, {
			Name:        []string{"--preview"},
			Description: `Command to preview highlighted line ({})`,
			Args: []model.Arg{{
				Name: "command",
			}},
		}, {
			Name:        []string{"--preview-window"},
			Description: `Preview window layout`,
			Args: []model.Arg{{
				Name:        "options",
				Suggestions: []model.Suggestion{{Name: []string{`up`}}, {Name: []string{`down`}}, {Name: []string{`left`}}, {Name: []string{`right`}}, {Name: []string{`nowrap`}}, {Name: []string{`wrap`}}, {Name: []string{`nocycle`}}, {Name: []string{`cycle`}}, {Name: []string{`nofollow`}}, {Name: []string{`follow`}}, {Name: []string{`nohidden`}}, {Name: []string{`hidden`}}, {Name: []string{`default`}}},
				IsVariadic:  true,
			}},
		}, {
			Name:        []string{"-q", "--query"},
			Description: `Start the finder with the given query`,
			Args: []model.Arg{{
				Name: "string",
			}},
		}, {
			Name:        []string{"-1", "--select-1"},
			Description: `Automatically select the only match`,
		}, {
			Name:        []string{"-0", "--exit-0"},
			Description: `Exit immediately when there's no match`,
		}, {
			Name:        []string{"-f", "--filter"},
			Description: `Filter mode. Do not start interactive finder`,
			Args: []model.Arg{{
				Name: "string",
			}},
		}, {
			Name:        []string{"--print-query"},
			Description: `Print query as the first line`,
		}, {
			Name:        []string{"--expect"},
			Description: `Comma-separated list of keys to complete fzf`,
			Args: []model.Arg{{
				Name: "keys",
			}},
		}, {
			Name:        []string{"--read0"},
			Description: `Read input delimited by ASCII NUL characters`,
		}, {
			Name:        []string{"--print0"},
			Description: `Print output delimited by ASCII NUL characters`,
		}, {
			Name:        []string{"--sync"},
			Description: `Synchronous search for multi-staged filtering`,
		}, {
			Name:        []string{"--version"},
			Description: `Display version information and exit`,
		}},
	}
}
