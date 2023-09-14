// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["ngrok"] = model.Subcommand{
		Name:        []string{"ngrok"},
		Description: `Tunnel local ports to public URLs and inspect traffic`,
		Subcommands: []model.Subcommand{{
			Name:        []string{"help"},
			Description: `Shows a list of commands or help for one command`,
			Args: []model.Arg{{
				Name: "command",
				Suggestions: []model.Suggestion{{
					Name:        []string{`authtoken`},
					Description: `Save authtoken to configuration file`,
				}, {
					Name:        []string{`credits`},
					Description: `Prints author and licensing information`,
				}, {
					Name:        []string{`http`},
					Description: `Start an HTTP tunnel`,
				}, {
					Name:        []string{`start`},
					Description: `Start tunnels by name from the configuration file`,
				}, {
					Name:        []string{`tcp`},
					Description: `Start a TCP tunnel`,
				}, {
					Name:        []string{`tls`},
					Description: `Start a TLS tunnel`,
				}, {
					Name:        []string{`update`},
					Description: `Update ngrok to the latest version`,
				}, {
					Name:        []string{`version`},
					Description: `Print the version string`,
				}, {
					Name:        []string{`help`},
					Description: `Shows a list of commands or help for one command`,
				}},
			}},
		}, {
			Name:        []string{"http"},
			Description: `Start an HTTP tunnel`,
			Args: []model.Arg{{
				Name: "host",
				Suggestions: []model.Suggestion{{
					Name:        []string{`8080`},
					Description: `Port`,
				}},
			}},
			Options: []model.Option{{
				Name:        []string{"--log", "-log"},
				Description: `Path to log file, 'stdout', 'stderr' or 'false'`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "value",
					Suggestions: []model.Suggestion{{Name: []string{`stdout`}}, {Name: []string{`stderr`}}, {Name: []string{`false`}}},
				}},
			}, {
				Name:        []string{"--log-format", "-log-format"},
				Description: `Log record format: 'term', 'logfmt', 'json'`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`term`}}, {Name: []string{`logfmt`}}, {Name: []string{`json`}}},
				}},
			}, {
				Name:        []string{"--log-level", "-log-level"},
				Description: `Logging level`,
				Args: []model.Arg{{
					Name:        "level",
					Suggestions: []model.Suggestion{{Name: []string{`crit`}}, {Name: []string{`warn`}}, {Name: []string{`error`}}, {Name: []string{`info`}}, {Name: []string{`debug`}}},
				}},
			}, {
				Name:        []string{"--config", "-config"},
				Description: `Path to config files; they are merged if multiple`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "config file",
				}},
			}, {
				Name:        []string{"--region", "-region"},
				Description: `Ngrok server region [us, eu, au, ap, sa, jp, in] (default: us)`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`us`}}, {Name: []string{`eu`}}, {Name: []string{`au`}}, {Name: []string{`ap`}}, {Name: []string{`sa`}}, {Name: []string{`jp`}}, {Name: []string{`in`}}},
				}},
			}, {
				Name:        []string{"--authtoken", "-authtoken"},
				Description: `Ngrok.com authtoken identifying a user`,
				Args: []model.Arg{{
					Name: "authtoken",
				}},
			}, {
				Name:        []string{"--hostname", "-hostname"},
				Description: `Host tunnel on custom hostname (requires DNS CNAME)`,
				Args: []model.Arg{{
					Name: "hostname",
				}},
			}, {
				Name:        []string{"--subdomain", "-subdomain"},
				Description: `Host tunnel on a custom subdomain`,
				Args: []model.Arg{{
					Name: "subdomain",
				}},
			}, {
				Name:        []string{"--auth", "-auth"},
				Description: `Enforce basic auth on tunnel endpoint, 'user:password'`,
				Args: []model.Arg{{
					Name: "user:password",
				}},
			}, {
				Name:        []string{"--bind-tls", "-bind-tls"},
				Description: `Listen for http, https or both: true/false/both`,
				Args: []model.Arg{{
					Name:        "true/false/both",
					Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}, {Name: []string{`both`}}},
				}},
			}, {
				Name:        []string{"--host-header", "-host-header"},
				Description: `Set Host header; if 'rewrite' use local address hostname`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`rewrite`}}},
				}},
			}, {
				Name:        []string{"--introspection", "-introspection"},
				Description: `Enable/disable http introspection`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}},
				}},
			}},
		}, {
			Name:        []string{"authtoken"},
			Description: `Save authtoken to configuration file`,
			Args: []model.Arg{{
				Name: "authtoken",
			}},
			Options: []model.Option{{
				Name:        []string{"--log", "-log"},
				Description: `Path to log file, 'stdout', 'stderr' or 'false'`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "value",
					Suggestions: []model.Suggestion{{Name: []string{`stdout`}}, {Name: []string{`stderr`}}, {Name: []string{`false`}}},
				}},
			}, {
				Name:        []string{"--log-format", "-log-format"},
				Description: `Log record format: 'term', 'logfmt', 'json'`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`term`}}, {Name: []string{`logfmt`}}, {Name: []string{`json`}}},
				}},
			}, {
				Name:        []string{"--log-level", "-log-level"},
				Description: `Logging level`,
				Args: []model.Arg{{
					Name:        "level",
					Suggestions: []model.Suggestion{{Name: []string{`crit`}}, {Name: []string{`warn`}}, {Name: []string{`error`}}, {Name: []string{`info`}}, {Name: []string{`debug`}}},
				}},
			}, {
				Name:        []string{"--config", "-config"},
				Description: `Path to config files; they are merged if multiple`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "config file",
				}},
			}},
		}, {
			Name:        []string{"credits"},
			Description: `Prints author and licensing information`,
		}, {
			Name:        []string{"start"},
			Description: `Start tunnels by name from the configuration file`,
			Args: []model.Arg{{
				Name:        "tunnels",
				Suggestions: []model.Suggestion{{Name: []string{`dev`}}, {Name: []string{`web`}}, {Name: []string{`blog`}}},
				IsVariadic:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"--log", "-log"},
				Description: `Path to log file, 'stdout', 'stderr' or 'false'`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "value",
					Suggestions: []model.Suggestion{{Name: []string{`stdout`}}, {Name: []string{`stderr`}}, {Name: []string{`false`}}},
				}},
			}, {
				Name:        []string{"--log-format", "-log-format"},
				Description: `Log record format: 'term', 'logfmt', 'json'`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`term`}}, {Name: []string{`logfmt`}}, {Name: []string{`json`}}},
				}},
			}, {
				Name:        []string{"--log-level", "-log-level"},
				Description: `Logging level`,
				Args: []model.Arg{{
					Name:        "level",
					Suggestions: []model.Suggestion{{Name: []string{`crit`}}, {Name: []string{`warn`}}, {Name: []string{`error`}}, {Name: []string{`info`}}, {Name: []string{`debug`}}},
				}},
			}, {
				Name:        []string{"--config", "-config"},
				Description: `Path to config files; they are merged if multiple`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "config file",
				}},
			}, {
				Name:        []string{"--region", "-region"},
				Description: `Ngrok server region [us, eu, au, ap, sa, jp, in] (default: us)`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`us`}}, {Name: []string{`eu`}}, {Name: []string{`au`}}, {Name: []string{`ap`}}, {Name: []string{`sa`}}, {Name: []string{`jp`}}, {Name: []string{`in`}}},
				}},
			}, {
				Name:        []string{"--authtoken", "-authtoken"},
				Description: `Ngrok.com authtoken identifying a user`,
				Args: []model.Arg{{
					Name: "authtoken",
				}},
			}, {
				Name:        []string{"--all", "-all"},
				Description: `Start all tunnels in the configuration file`,
				ExclusiveOn: []string{"--none"},
			}, {
				Name:        []string{"--none", "-none"},
				Description: `Start running no tunnels`,
				ExclusiveOn: []string{"--all"},
			}},
		}, {
			Name:        []string{"tcp"},
			Description: `Start a TCP tunnel`,
			Args: []model.Arg{{
				Name:        "port",
				Suggestions: []model.Suggestion{{Name: []string{`22`}}},
			}},
			Options: []model.Option{{
				Name:        []string{"--log", "-log"},
				Description: `Path to log file, 'stdout', 'stderr' or 'false'`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "value",
					Suggestions: []model.Suggestion{{Name: []string{`stdout`}}, {Name: []string{`stderr`}}, {Name: []string{`false`}}},
				}},
			}, {
				Name:        []string{"--log-format", "-log-format"},
				Description: `Log record format: 'term', 'logfmt', 'json'`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`term`}}, {Name: []string{`logfmt`}}, {Name: []string{`json`}}},
				}},
			}, {
				Name:        []string{"--log-level", "-log-level"},
				Description: `Logging level`,
				Args: []model.Arg{{
					Name:        "level",
					Suggestions: []model.Suggestion{{Name: []string{`crit`}}, {Name: []string{`warn`}}, {Name: []string{`error`}}, {Name: []string{`info`}}, {Name: []string{`debug`}}},
				}},
			}, {
				Name:        []string{"--config", "-config"},
				Description: `Path to config files; they are merged if multiple`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "config file",
				}},
			}, {
				Name:        []string{"--authtoken", "-authtoken"},
				Description: `Ngrok.com authtoken identifying a user`,
				Args: []model.Arg{{
					Name: "authtoken",
				}},
			}, {
				Name:        []string{"--region", "-region"},
				Description: `Ngrok server region [us, eu, au, ap, sa, jp, in] (default: us)`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`us`}}, {Name: []string{`eu`}}, {Name: []string{`au`}}, {Name: []string{`ap`}}, {Name: []string{`sa`}}, {Name: []string{`jp`}}, {Name: []string{`in`}}},
				}},
			}, {
				Name:        []string{"--remote-addr", "-remote-addr"},
				Description: `Bind remote address (requires you reserve an address)`,
				Args: []model.Arg{{
					Name: "remote address",
				}},
			}},
		}, {
			Name:        []string{"tls"},
			Description: `Start a TLS tunnel`,
			Args: []model.Arg{{
				Name: "port",
			}},
			Options: []model.Option{{
				Name:        []string{"--log", "-log"},
				Description: `Path to log file, 'stdout', 'stderr' or 'false'`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "value",
					Suggestions: []model.Suggestion{{Name: []string{`stdout`}}, {Name: []string{`stderr`}}, {Name: []string{`false`}}},
				}},
			}, {
				Name:        []string{"--log-format", "-log-format"},
				Description: `Log record format: 'term', 'logfmt', 'json'`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`term`}}, {Name: []string{`logfmt`}}, {Name: []string{`json`}}},
				}},
			}, {
				Name:        []string{"--log-level", "-log-level"},
				Description: `Logging level`,
				Args: []model.Arg{{
					Name:        "level",
					Suggestions: []model.Suggestion{{Name: []string{`crit`}}, {Name: []string{`warn`}}, {Name: []string{`error`}}, {Name: []string{`info`}}, {Name: []string{`debug`}}},
				}},
			}, {
				Name:        []string{"--config", "-config"},
				Description: `Path to config files; they are merged if multiple`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "config file",
				}},
			}, {
				Name:        []string{"--authtoken", "-authtoken"},
				Description: `Ngrok.com authtoken identifying a user`,
				Args: []model.Arg{{
					Name: "authtoken",
				}},
			}, {
				Name:        []string{"--region", "-region"},
				Description: `Ngrok server region [us, eu, au, ap, sa, jp, in] (default: us)`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`us`}}, {Name: []string{`eu`}}, {Name: []string{`au`}}, {Name: []string{`ap`}}, {Name: []string{`sa`}}, {Name: []string{`jp`}}, {Name: []string{`in`}}},
				}},
			}, {
				Name:        []string{"--hostname", "-hostname"},
				Description: `Host tunnel on custom hostname (requires DNS CNAME)`,
				Args: []model.Arg{{
					Name: "hostname",
				}},
			}, {
				Name:        []string{"--subdomain", "-subdomain"},
				Description: `Host tunnel on a custom subdomain`,
				Args: []model.Arg{{
					Name: "subdomain",
				}},
			}, {
				Name: []string{"--client-cas", "-client-cas"},
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "certificate",
				}},
			}, {
				Name: []string{"--crt", "-crt"},
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "certificate",
				}},
			}, {
				Name: []string{"--key", "-key"},
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "certificate",
				}},
			}},
		}, {
			Name:        []string{"update"},
			Description: `Update ngrok to the latest version`,
			Options: []model.Option{{
				Name:        []string{"--log", "-log"},
				Description: `Path to log file, 'stdout', 'stderr' or 'false'`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "value",
					Suggestions: []model.Suggestion{{Name: []string{`stdout`}}, {Name: []string{`stderr`}}, {Name: []string{`false`}}},
				}},
			}, {
				Name:        []string{"--log-format", "-log-format"},
				Description: `Log record format: 'term', 'logfmt', 'json'`,
				Args: []model.Arg{{
					Suggestions: []model.Suggestion{{Name: []string{`term`}}, {Name: []string{`logfmt`}}, {Name: []string{`json`}}},
				}},
			}, {
				Name:        []string{"--log-level", "-log-level"},
				Description: `Logging level`,
				Args: []model.Arg{{
					Name:        "level",
					Suggestions: []model.Suggestion{{Name: []string{`crit`}}, {Name: []string{`warn`}}, {Name: []string{`error`}}, {Name: []string{`info`}}, {Name: []string{`debug`}}},
				}},
			}, {
				Name:        []string{"--channel", "-channel"},
				Description: `Update channel (stable, beta)`,
				Args: []model.Arg{{
					Name:        "channel",
					Suggestions: []model.Suggestion{{Name: []string{`stable`}}, {Name: []string{`beta`}}},
				}},
			}},
		}, {
			Name:        []string{"version"},
			Description: `Print the version string`,
		}},
	}
}
