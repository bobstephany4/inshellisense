// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["http"] = model.Subcommand{
		Name:        []string{"http"},
		Description: `HTTPie: command-line HTTP client for the API era`,
		Args: []model.Arg{{
			Name:        "METHOD",
			Description: `The HTTP method to be used for the request (GET, POST, PUT, DELETE, ...)`,
			Suggestions: []model.Suggestion{{Name: []string{`GET`}}, {Name: []string{`POST`}}, {Name: []string{`PUT`}}, {Name: []string{`DELETE`}}, {Name: []string{`HEAD`}}},
			IsOptional:  true,
		}, {
			Name:        "URL",
			Description: `The scheme defaults to 'http://' if the URL does not include one`,
		}, {
			Name:        "REQUEST_ITEM",
			Description: `Optional key-value pairs to be included in the request. The separator used determines the type`,
			IsOptional:  true,
			IsVariadic:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"--json", "-j"},
			Description: `Data items from the command line are serialized as a JSON object. The Content-Type and Accept headers are set to application/json`,
			ExclusiveOn: []string{"--form", "-f", "--multipart", "--boundary"},
		}, {
			Name: []string{"--form", "-f"},
			Description: `FData items from the command line are serialized as form fields.
The Content-Type is set to application/x-www-form-urlencoded (if not
specified). The presence of any file fields results in a
multipart/form-data request`,
			ExclusiveOn: []string{"--json", "-j", "--multipart", "--boundary"},
		}, {
			Name:        []string{"--multipart"},
			Description: `Similar to --form, but always sends a multipart/form-data request (i.e., even without files)`,
			ExclusiveOn: []string{"--json", "-j", "--form", "-f", "--boundary"},
		}, {
			Name:        []string{"--boundary"},
			Description: `Specify a custom boundary string for multipart/form-data requests. Has effect only together with --form`,
			ExclusiveOn: []string{"--json", "-j", "--multipart"},
		}, {
			Name:        []string{"--compress", "-x"},
			Description: `Content compressed (encoded) with Deflate algorithm. The Content-Encoding header is set to deflate`,
		}, {
			Name: []string{"--pretty"},
			Description: `Controls output processing. The value can be "none" to not prettify
the output (default for redirected output), "all" to apply both colors
and formatting (default for terminal output), "colors", or "format"`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{Name: []string{`all`}}, {Name: []string{`colors`}}, {Name: []string{`format`}}, {Name: []string{`none`}}},
			}},
		}, {
			Name: []string{"--style", "-s"},
			Args: []model.Arg{{
				Name:        "STYLE",
				Suggestions: []model.Suggestion{{Name: []string{`abap`}}, {Name: []string{`algol`}}, {Name: []string{`algol_nu`}}, {Name: []string{`arduino`}}, {Name: []string{`auto`}}, {Name: []string{`autumn`}}, {Name: []string{`borland`}}, {Name: []string{`bw`}}, {Name: []string{`colorful`}}, {Name: []string{`default`}}, {Name: []string{`emacs`}}, {Name: []string{`friendly`}}, {Name: []string{`fruity`}}, {Name: []string{`igor`}}, {Name: []string{`inkpot`}}, {Name: []string{`lovelace`}}, {Name: []string{`manni`}}, {Name: []string{`monokai`}}, {Name: []string{`murphy`}}, {Name: []string{`native`}}, {Name: []string{`paraiso-dark`}}, {Name: []string{`paraiso-light`}}, {Name: []string{`pastie`}}, {Name: []string{`perldoc`}}, {Name: []string{`rainbow_dash`}}, {Name: []string{`rrt`}}, {Name: []string{`sas`}}, {Name: []string{`solarized`}}, {Name: []string{`solarized-dark`}}, {Name: []string{`solarized-light`}}, {Name: []string{`stata`}}, {Name: []string{`stata-dark`}}, {Name: []string{`stata-light`}}, {Name: []string{`tango`}}, {Name: []string{`trac`}}, {Name: []string{`vim`}}, {Name: []string{`vs`}}, {Name: []string{`xcode`}}},
			}},
		}, {
			Name:        []string{"--unsorted"},
			Description: `Disables all sorting while formatting output`,
			ExclusiveOn: []string{"--sorted"},
		}, {
			Name:        []string{"--sorted"},
			Description: `Re-enables all sorting options while formatting output`,
			ExclusiveOn: []string{"--unsorted"},
		}, {
			Name:        []string{"--format-options"},
			Description: `Controls output formatting`,
			Args: []model.Arg{{
				Name:        "FORMAT_OPTIONS",
				Suggestions: []model.Suggestion{{Name: []string{`headers.sort:true`}}, {Name: []string{`json.format:true`}}, {Name: []string{`json.indent:4`}}, {Name: []string{`json.sort_keys:true`}}},
			}},
		}, {
			Name:        []string{"--print", "-p"},
			Description: `String specifying what the output should contain`,
			Args: []model.Arg{{
				Name: "WHAT",
				Suggestions: []model.Suggestion{{
					Name:        []string{`H`},
					Description: `Request headers`,
				}, {
					Name:        []string{`B`},
					Description: `Request body`,
				}, {
					Name:        []string{`h`},
					Description: `Response headers`,
				}, {
					Name:        []string{`b`},
					Description: `Response body`,
				}},
			}},
		}, {
			Name:        []string{"--headers", "-h"},
			Description: `Print only the response headers. Shortcut for --print=h`,
		}, {
			Name:        []string{"--body", "-b"},
			Description: `Print only the response body. Shortcut for --print=b`,
		}, {
			Name: []string{"--verbose", "-v"},
			Description: `Verbose output. Print the whole request as well as the response. Also print
any intermediary requests/responses (such as redirects).
It's a shortcut for: --all --print=BHbh`,
		}, {
			Name: []string{"--all"},
			Description: `By default, only the final request/response is shown. Use this flag to show
any intermediary requests/responses as well. Intermediary requests include
followed redirects (with --follow), the first unauthorized request when
Digest auth is used (--auth=digest), etc`,
		}, {
			Name: []string{"--history-print", "-P"},
			Description: `The same as --print, -p but applies only to intermediary requests/responses
(such as redirects) when their inclusion is enabled with --all. If this
options is not specified, then they are formatted the same way as the final
response`,
			Args: []model.Arg{{
				Name: "WHAT",
				Suggestions: []model.Suggestion{{
					Name:        []string{`H`},
					Description: `Request headers`,
				}, {
					Name:        []string{`B`},
					Description: `Request body`,
				}, {
					Name:        []string{`h`},
					Description: `Response headers`,
				}, {
					Name:        []string{`b`},
					Description: `Response body`,
				}},
			}},
		}, {
			Name:        []string{"--stream", "-S"},
			Description: `Always stream the response body by line, i.e., behave like "tail -f'`,
		}, {
			Name: []string{"--output", "-o"},
			Description: `Save output to FILE instead of stdout. If --download is also set, then only
the response body is saved to FILE. Other parts of the HTTP exchange are
printed to stderr`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "FILE",
			}},
		}, {
			Name: []string{"--download", "-d"},
			Description: `Do not print the response body to stdout. Rather, download it and store it
in a file. The filename is guessed unless specified with --output
[filename]. This action is similar to the default behaviour of wget`,
		}, {
			Name:        []string{"--continue", "-c"},
			Description: `Resume an interrupted download. Note that the --output option needs to be specified as well`,
		}, {
			Name: []string{"--quiet", "-q"},
			Description: `Do not print to stdout or stderr.
stdout is still redirected if --output is specified.
Flag doesn't affect behaviour of download beyond not printing to terminal`,
		}, {
			Name: []string{"--session"},
			Description: `Create, or reuse and update a session. Within a session, custom headers,
auth credential, as well as any cookies sent by the server persist between
requests`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "SESSION_NAME_OR_PATH",
			}},
		}, {
			Name:        []string{"--session-read-only"},
			Description: `Create or read a session without updating it form the request/response exchange`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "SESSION_NAME_OR_PATH",
			}},
		}, {
			Name:        []string{"--auth", "-a"},
			Description: `If only the username is provided (-a username), HTTPie will prompt for the password`,
			Args: []model.Arg{{
				Name: "USER[:PASS]",
			}},
		}, {
			Name:        []string{"--auth-type", "-A"},
			Description: `The authentication mechanism to be used. Defaults to "basic"`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{
					Name:        []string{`basic`},
					Description: `Basic HTTP auth`,
				}, {
					Name:        []string{`digest`},
					Description: `Digest HTTP auth`,
				}},
			}},
		}, {
			Name:        []string{"--ignore-netrc"},
			Description: `Ignore credentials from .netrc`,
		}, {
			Name:        []string{"--offline"},
			Description: `Build the request and print it but don’t actually send it`,
		}, {
			Name: []string{"--proxy"},
			Description: `String mapping protocol to the URL of the proxy
(e.g. http:http://foo.bar:3128). You can specify multiple proxies with
different protocols. The environment variables $ALL_PROXY, $HTTP_PROXY,
and $HTTPS_proxy are supported as well`,
			Args: []model.Arg{{
				Name: "PROTOCOL:PROXY_URL",
			}},
		}, {
			Name:        []string{"--follow", "-F"},
			Description: `Follow 30x Location redirects`,
		}, {
			Name:        []string{"--max-redirects"},
			Description: `By default, requests have a limit of 30 redirects (works with --follow)`,
			Args: []model.Arg{{
				Name: "MAX_REDIRECTS",
			}},
		}, {
			Name:        []string{"--max-headers"},
			Description: `The maximum number of response headers to be read before giving up (default 0, i.e., no limit)`,
			Args: []model.Arg{{
				Name: "MAX_HEADERS",
			}},
		}, {
			Name: []string{"--timeout"},
			Description: `The connection timeout of the request in seconds.
The default value is 0, i.e., there is no timeout limit`,
			Args: []model.Arg{{
				Name: "SECONDS",
			}},
		}, {
			Name: []string{"--check-status"},
			Description: `By default, HTTPie exits with 0 when no network or other fatal errors
occur. This flag instructs HTTPie to also check the HTTP status code and
exit with an error if the status indicates one`,
		}, {
			Name:        []string{"--path-as-is"},
			Description: `Bypass dot segment (/../ or /./) URL squashing`,
		}, {
			Name:        []string{"--chunked"},
			Description: `Enable streaming via chunked transfer encoding`,
		}, {
			Name: []string{"--verify"},
			Description: `Set to "no" (or "false") to skip checking the host's SSL certificate.
Defaults to "yes" ("true"). You can also pass the path to a CA_BUNDLE file
for private certs`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFilepaths},
				Name:        "VERIFY",
				Suggestions: []model.Suggestion{{Name: []string{`no`}}, {Name: []string{`yes`}}},
			}},
		}, {
			Name: []string{"--ssl"},
			Description: `The desired protocol version to use. This will default to
SSL v2.3 which will negotiate the highest protocol that both
the server and your installation of OpenSSL support`,
			Args: []model.Arg{{
				Suggestions: []model.Suggestion{{Name: []string{`ssl2.3`}}, {Name: []string{`tls1`}}, {Name: []string{`tls1.1`}}, {Name: []string{`tls1.2`}}},
			}},
		}, {
			Name:        []string{"--ciphers"},
			Description: `A string in the OpenSSL cipher list format`,
			Args:        []model.Arg{{}},
		}, {
			Name: []string{"--cert"},
			Description: `You can specify a local cert to use as client side SSL certificate.
This file may either contain both private key and certificate or you may
specify --cert-key separately`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
			}},
		}, {
			Name:        []string{"--cert-key"},
			Description: `The private key to use with SSL. Only needed if --cert is given and the certificate file does not contain the private key`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
			}},
		}, {
			Name:        []string{"--ignore-stdin", "-I"},
			Description: `Do not attempt to read stdin`,
		}, {
			Name:        []string{"--help"},
			Description: `Show the help message and exit`,
		}, {
			Name:        []string{"--version"},
			Description: `Show version and exit`,
		}, {
			Name:        []string{"--traceback"},
			Description: `Prints the exception traceback should one occur`,
		}, {
			Name:        []string{"--default-scheme"},
			Description: `The default scheme to use if not specified in the URL`,
			Args: []model.Arg{{
				Name: "DEFAULT_SCHEME",
			}},
		}, {
			Name: []string{"--debug"},
			Description: `Prints the exception traceback should one occur, as well as other
information useful for debugging HTTPie itself and for reporting bugs`,
		}},
	}
}
