// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["screen"] = model.Subcommand{
		Name:        []string{"screen"},
		Description: `Screen manager with VT100/ANSI terminal emulation`,
		Args: []model.Arg{{
			Name:        "cmd",
			Description: `Command to invoke`,
		}, {
			Name:        "args",
			Description: `Arguments to pass`,
			IsVariadic:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"-a"},
			Description: `Force all capabilities into each window's termcap`,
		}, {
			Name:        []string{"-A"},
			Description: `Adapt all windows to the new display width & height`,
		}, {
			Name:        []string{"-c"},
			Description: `Read configuration file instead of '.screenrc'`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "file",
			}},
		}, {
			Name:        []string{"-e"},
			Description: `Change command characters`,
			Args: []model.Arg{{
				Name:        "xy",
				Description: `Specifies the command character to be x and the character generating a literal command character to y`,
			}},
		}, {
			Name:        []string{"-f"},
			Description: `Flow control on`,
		}, {
			Name:        []string{"-fn"},
			Description: `Flow control off`,
		}, {
			Name:        []string{"-fa"},
			Description: `Flow control automatic`,
		}, {
			Name:        []string{"-h"},
			Description: `Set the size of the scrollback history buffer`,
			Args: []model.Arg{{
				Name:        "h",
				Description: `Lines`,
			}},
		}, {
			Name:        []string{"-i"},
			Description: `Interrupt output sooner when flow control is on`,
		}, {
			Name:        []string{"-list", "-ls"},
			Description: `Do nothing, just list our SockDir`,
		}, {
			Name:        []string{"-L"},
			Description: `Turn on output logging`,
		}, {
			Name:        []string{"-m"},
			Description: `Ignore $STY variable, do create a new screen session`,
		}, {
			Name:        []string{"-O"},
			Description: `Choose optimal output rather than exact vt100 emulation`,
		}, {
			Name:        []string{"-p"},
			Description: `Preselect the named window if it exists`,
			Args: []model.Arg{{
				Name: "window",
			}},
		}, {
			Name:        []string{"-q"},
			Description: `Quiet startup. Exits with non-zero return code if unsuccessful`,
		}, {
			Name:        []string{"-s"},
			Description: `Shell to execute rather than $SHELL`,
			Args: []model.Arg{{
				Name: "shell",
			}},
		}, {
			Name:        []string{"-S"},
			Description: `Name this session <pid>.sockname instead of <pid>.<tty>.<host>`,
			Args: []model.Arg{{
				Name: "sockname",
			}},
		}, {
			Name:        []string{"-t"},
			Description: `Set title. (window's name)`,
			Args: []model.Arg{{
				Name: "title",
			}},
		}, {
			Name:        []string{"-T"},
			Description: `Use term as $TERM for windows, rather than 'screen'`,
			Args: []model.Arg{{
				Name: "term",
			}},
		}, {
			Name:        []string{"-U"},
			Description: `Tell screen to use UTF-8 encoding`,
		}, {
			Name:        []string{"-v"},
			Description: `Print 'Screen version 4.00.03 (FAU) 23-Oct-06'`,
		}, {
			Name:        []string{"-wipe"},
			Description: `Do nothing, just clean up SockDir`,
			Args: []model.Arg{{
				Name:        "match",
				Description: `Screen session to match`,
				Suggestions: []model.Suggestion{{Name: []string{`sessionowner/[pid.tty.host]`}}},
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"-x"},
			Description: `Attach to a not detached screen. (Multi display mode)`,
		}, {
			Name:        []string{"-X"},
			Description: `Execute <cmd> as a screen command in the specified session`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"-d"},
			Description: `Does not start screen, but detaches the elsewhere running screen session`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"-r"},
				Description: `Reattach a session and if necessary detach it first`,
				Args: []model.Arg{{
					Name:        "identifier",
					Suggestions: []model.Suggestion{{Name: []string{`pid`}}, {Name: []string{`pid.tty`}}, {Name: []string{`pid.tty.host`}}, {Name: []string{`sessionowner/pid`}}, {Name: []string{`sessionowner/pid.tty`}}, {Name: []string{`sessionowner/pid.tty.host`}}},
				}},
			}, {
				Name:        []string{"-R"},
				Description: `Reattach a session and if necessary detach or even create it first`,
				Args: []model.Arg{{
					Name:        "identifier",
					Suggestions: []model.Suggestion{{Name: []string{`pid`}}, {Name: []string{`pid.tty`}}, {Name: []string{`pid.tty.host`}}, {Name: []string{`sessionowner/pid`}}, {Name: []string{`sessionowner/pid.tty`}}, {Name: []string{`sessionowner/pid.tty.host`}}},
				}},
			}, {
				Name:        []string{"-RR"},
				Description: `Reattach a session and if necessary detach or create it. Use the first session if more than one session is available`,
				Args: []model.Arg{{
					Name:        "identifier",
					Suggestions: []model.Suggestion{{Name: []string{`pid`}}, {Name: []string{`pid.tty`}}, {Name: []string{`pid.tty.host`}}, {Name: []string{`sessionowner/pid`}}, {Name: []string{`sessionowner/pid.tty`}}, {Name: []string{`sessionowner/pid.tty.host`}}},
				}},
			}},
		}, {
			Name:        []string{"-D"},
			Description: `Does not start screen, but detaches the elsewhere running screen session`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"-r"},
				Description: `Reattach a session. If necessary detach and logout remotely first`,
				Args: []model.Arg{{
					Name:        "identifier",
					Suggestions: []model.Suggestion{{Name: []string{`pid`}}, {Name: []string{`pid.tty`}}, {Name: []string{`pid.tty.host`}}, {Name: []string{`sessionowner/pid`}}, {Name: []string{`sessionowner/pid.tty`}}, {Name: []string{`sessionowner/pid.tty.host`}}},
				}},
			}, {
				Name:        []string{"-R"},
				Description: `Attach here and now. In detail this means: If a session is running, then reattach. If necessary detach and logout remotely first`,
				Args: []model.Arg{{
					Name:        "identifier",
					Suggestions: []model.Suggestion{{Name: []string{`pid`}}, {Name: []string{`pid.tty`}}, {Name: []string{`pid.tty.host`}}, {Name: []string{`sessionowner/pid`}}, {Name: []string{`sessionowner/pid.tty`}}, {Name: []string{`sessionowner/pid.tty.host`}}},
				}},
			}, {
				Name:        []string{"-RR"},
				Description: `Attach here and now. Whatever that means, just do it`,
				Args: []model.Arg{{
					Name:        "identifier",
					Suggestions: []model.Suggestion{{Name: []string{`pid`}}, {Name: []string{`pid.tty`}}, {Name: []string{`pid.tty.host`}}, {Name: []string{`sessionowner/pid`}}, {Name: []string{`sessionowner/pid.tty`}}, {Name: []string{`sessionowner/pid.tty.host`}}},
				}},
			}},
		}, {
			Name:        []string{"-r"},
			Description: `Resumes a detached screen session`,
			Args: []model.Arg{{
				Name:        "identifier",
				Suggestions: []model.Suggestion{{Name: []string{`pid`}}, {Name: []string{`pid.tty`}}, {Name: []string{`pid.tty.host`}}, {Name: []string{`sessionowner/pid`}}, {Name: []string{`sessionowner/pid.tty`}}, {Name: []string{`sessionowner/pid.tty.host`}}},
			}},
		}, {
			Name:        []string{"-R"},
			Description: `Attempts to resume the first detached screen session it finds`,
			Args: []model.Arg{{
				Name:        "identifier",
				Suggestions: []model.Suggestion{{Name: []string{`pid`}}, {Name: []string{`pid.tty`}}, {Name: []string{`pid.tty.host`}}, {Name: []string{`sessionowner/pid`}}, {Name: []string{`sessionowner/pid.tty`}}, {Name: []string{`sessionowner/pid.tty.host`}}},
			}},
		}, {
			Name:        []string{"-dmS"},
			Description: `Start as daemon: Screen session in detached mode`,
			Args: []model.Arg{{
				Name:        "name",
				Description: `Name of the screen session`,
			}},
		}},
	}
}
