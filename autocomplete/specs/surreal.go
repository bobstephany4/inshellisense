// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["surreal"] = model.Subcommand{
		Name:        []string{"surreal"},
		Description: `SurrealDB is the ultimate cloud database for tomorrow's applications - https://surrealdb.com/`,
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Print help information`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"help"},
			Description: `Print this message or the help of the given subcommand(s)`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateHelp},
				Name:        "command",
				Description: `Command to get help for`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"start"},
			Description: `Start the database server`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Database path used for storing data [env: DB_PATH=] [default: memory]`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"--"},
				Description: `Everything after this is an argument`,
			}, {
				Name:        []string{"--addr"},
				Description: `The allowed networks for master authentication [env: ADDR=] [default: 127.0.0.1/32]`,
				Args: []model.Arg{{
					Name:        "addr",
					Description: `The allowed networks for master authentication [env: ADDR=] [default: 127.0.0.1/32]`,
				}},
			}, {
				Name:        []string{"--bind", "-b"},
				Description: `The hostname or ip address to listen for connections on [env: BIND=] [default: 0.0.0.0:8000]`,
				Args: []model.Arg{{
					Name:        "bind",
					Description: `The hostname or ip address to listen for connections on [env: BIND=] [default: 0.0.0.0:8000]`,
				}},
			}, {
				Name:        []string{"--key", "-k"},
				Description: `Encryption key to use for on-disk encryption [env: KEY=]`,
				Args: []model.Arg{{
					Name:        "key",
					Description: `Encryption key to use for on-disk encryption [env: KEY=]`,
				}},
			}, {
				Name:        []string{"--kvs-ca"},
				Description: `Path to the CA file used when connecting to the remote KV store [env: KVS_CA=]`,
				Args: []model.Arg{{
					Name:        "kvs-ca",
					Description: `Path to the CA file used when connecting to the remote KV store [env: KVS_CA=]`,
				}},
			}, {
				Name:        []string{"--kvs-crt"},
				Description: `Path to the certificate file used when connecting to the remote KV store [env: KVS_CRT=]`,
				Args: []model.Arg{{
					Name:        "kvs-crt",
					Description: `Path to the certificate file used when connecting to the remote KV store [env: KVS_CRT=]`,
				}},
			}, {
				Name:        []string{"--kvs-key"},
				Description: `Path to the private key file used when connecting to the remote KV store [env: KVS_KEY=]`,
				Args: []model.Arg{{
					Name:        "kvs-key",
					Description: `Path to the private key file used when connecting to the remote KV store [env: KVS_KEY=]`,
				}},
			}, {
				Name:        []string{"--log", "-l"},
				Description: `The logging level for the database server [env: LOG=] [default: info] [possible values: warn, info, debug, trace, full]`,
				Args: []model.Arg{{
					Name:        "log",
					Description: `The logging level for the database server [env: LOG=] [default: info] [possible values: warn, info, debug, trace, full]`,
					Suggestions: []model.Suggestion{{Name: []string{`warn`}}, {Name: []string{`info`}}, {Name: []string{`debug`}}, {Name: []string{`trace`}}, {Name: []string{`full`}}},
				}},
			}, {
				Name:        []string{"--pass", "-p"},
				Description: `The master password for the database [env: PASS=]`,
				Args: []model.Arg{{
					Name:        "pass",
					Description: `The master password for the database [env: PASS=]`,
				}},
			}, {
				Name:        []string{"--strict", "-s"},
				Description: `Whether strict mode is enabled on this database instance [env: STRICT=]`,
				Args: []model.Arg{{
					Name:        "strict",
					Description: `Whether strict mode is enabled on this database instance [env: STRICT=]`,
				}},
			}, {
				Name:        []string{"--user", "-u"},
				Description: `The master username for the database [env: USER=] [default: root]`,
				Args: []model.Arg{{
					Name:        "user",
					Description: `The master username for the database [env: USER=] [default: root]`,
				}},
			}, {
				Name:        []string{"--web-crt"},
				Description: `Path to the certificate file for encrypted client connections [env: WEB_CRT=]`,
				Args: []model.Arg{{
					Name:        "web-crt",
					Description: `Path to the certificate file for encrypted client connections [env: WEB_CRT=]`,
				}},
			}, {
				Name:        []string{"--web-key"},
				Description: `Path to the private key file for encrypted client connections [env: WEB_KEY=]`,
				Args: []model.Arg{{
					Name:        "web-key",
					Description: `Path to the private key file for encrypted client connections [env: WEB_KEY=]`,
				}},
			}},
		}, {
			Name:        []string{"backup"},
			Description: `Backup data to or from an existing database`,
			Args: []model.Arg{{
				Name:        "from",
				Description: `Path to the remote database or file from which to export`,
			}, {
				Name:        "into",
				Description: `Path to the remote database or file into which to import`,
			}},
			Options: []model.Option{{
				Name:        []string{"--pass", "-p"},
				Description: `Database authentication password to use when connecting [default: root]`,
				Args: []model.Arg{{
					Name:        "pass",
					Description: `Database authentication password to use when connecting [default: root]`,
				}},
			}, {
				Name:        []string{"--user", "-u"},
				Description: `Database authentication username to use when connecting [default: root]`,
				Args: []model.Arg{{
					Name:        "user",
					Description: `Database authentication username to use when connecting [default: root]`,
				}},
			}},
		}, {
			Name:        []string{"import"},
			Description: `Import a SurrealQL script into an existing database`,
			Args: []model.Arg{{
				Name:        "file",
				Description: `Path to the sql file to import`,
			}},
			Options: []model.Option{{
				Name:        []string{"--pass", "-p"},
				Description: `Database authentication password to use when connecting [default: root]`,
				Args: []model.Arg{{
					Name:        "pass",
					Description: `Database authentication password to use when connecting [default: root]`,
				}},
			}, {
				Name:        []string{"--user", "-u"},
				Description: `Database authentication username to use when connecting [default: root]`,
				Args: []model.Arg{{
					Name:        "user",
					Description: `Database authentication username to use when connecting [default: root]`,
				}},
			}, {
				Name:        []string{"--conn", "-c"},
				Description: `Remote database server url to connect to [default: https://cloud.surrealdb.com]`,
				Args: []model.Arg{{
					Name:        "conn",
					Description: `Remote database server url to connect to [default: https://cloud.surrealdb.com]`,
				}},
			}, {
				Name:        []string{"--ns"},
				Description: `The namespace to import the data into`,
				Args: []model.Arg{{
					Name:        "ns",
					Description: `The namespace to import the data into`,
				}},
			}, {
				Name:        []string{"--db"},
				Description: `The database to import the data into`,
				Args: []model.Arg{{
					Name:        "db",
					Description: `The database to import the data into`,
				}},
			}},
		}, {
			Name:        []string{"export"},
			Description: `Export an existing database as a SurrealQL script`,
			Args: []model.Arg{{
				Name:        "file",
				Description: `Path to the sql file to export`,
			}},
			Options: []model.Option{{
				Name:        []string{"--pass", "-p"},
				Description: `Database authentication password to use when connecting [default: root]`,
				Args: []model.Arg{{
					Name:        "pass",
					Description: `Database authentication password to use when connecting [default: root]`,
				}},
			}, {
				Name:        []string{"--user", "-u"},
				Description: `Database authentication username to use when connecting [default: root]`,
				Args: []model.Arg{{
					Name:        "user",
					Description: `Database authentication username to use when connecting [default: root]`,
				}},
			}, {
				Name:        []string{"--conn", "-c"},
				Description: `Remote database server url to connect to [default: https://cloud.surrealdb.com]`,
				Args: []model.Arg{{
					Name:        "conn",
					Description: `Remote database server url to connect to [default: https://cloud.surrealdb.com]`,
				}},
			}, {
				Name:        []string{"--ns"},
				Description: `The namespace to export the data from`,
				Args: []model.Arg{{
					Name:        "ns",
					Description: `The namespace to export the data from`,
				}},
			}, {
				Name:        []string{"--db"},
				Description: `The database to export the data from`,
				Args: []model.Arg{{
					Name:        "db",
					Description: `The database to export the data from`,
				}},
			}},
		}, {
			Name:        []string{"version"},
			Description: `Output the command-line tool version information`,
		}, {
			Name:        []string{"sql"},
			Description: `Start an SQL REPL in your terminal with pipe support`,
			Args: []model.Arg{{
				Name:        "file",
				Description: `Path to the sql file to export`,
			}},
			Options: []model.Option{{
				Name:        []string{"--pass", "-p"},
				Description: `Database authentication password to use when connecting [default: root]`,
				Args: []model.Arg{{
					Name:        "pass",
					Description: `Database authentication password to use when connecting [default: root]`,
				}},
			}, {
				Name:        []string{"--user", "-u"},
				Description: `Database authentication username to use when connecting [default: root]`,
				Args: []model.Arg{{
					Name:        "user",
					Description: `Database authentication username to use when connecting [default: root]`,
				}},
			}, {
				Name:        []string{"--conn", "-c"},
				Description: `Remote database server url to connect to [default: https://cloud.surrealdb.com]`,
				Args: []model.Arg{{
					Name:        "conn",
					Description: `Remote database server url to connect to [default: https://cloud.surrealdb.com]`,
				}},
			}, {
				Name:        []string{"--ns"},
				Description: `The namespace to export the data from`,
				Args: []model.Arg{{
					Name:        "ns",
					Description: `The namespace to export the data from`,
				}},
			}, {
				Name:        []string{"--db"},
				Description: `The database to export the data from`,
				Args: []model.Arg{{
					Name:        "db",
					Description: `The database to export the data from`,
				}},
			}, {
				Name:        []string{"--pretty"},
				Description: `Whether database responses should be pretty printed`,
			}},
		}},
	}
}
