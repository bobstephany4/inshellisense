// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["meteor"] = model.Subcommand{
		Name:        []string{"meteor"},
		Description: `Run the meteor command-line tool`,
		Subcommands: []model.Subcommand{{
			Name: []string{"npm"},
		}, {
			Name:        []string{"help"},
			Description: `Get help on meteor command line usage`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateHelp},
				Name:        "command",
				Description: `Built-in command`,
				IsOptional:  true,
			}},
			Subcommands: []model.Subcommand{{
				Name: []string{"admin"},
				Subcommands: []model.Subcommand{{
					Name: []string{"maintainers"},
				}, {
					Name: []string{"recommend-release"},
				}, {
					Name: []string{"change-homepage"},
				}, {
					Name: []string{"list-organizations"},
				}, {
					Name: []string{"members"},
				}},
			}},
		}, {
			Name:        []string{"run"},
			Description: `Run a meteor development server in the current project. Searches upward from the current directory for the root directory of a Meteor project`,
			Options: []model.Option{{
				Name:        []string{"target"},
				Description: `If you have added a platform to your app with 'meteor add-platform', you can pass one of the following targets as an argument to this command`,
			}, {
				Name:        []string{"--port", "-p"},
				Description: `Port to listen on (instead of the default 3000)`,
				Args: []model.Arg{{
					Name:       "port",
					IsOptional: true,
				}},
			}, {
				Name:        []string{"--inspect"},
				Description: `Enable server-side debugging via debugging clients like the Node.js command-line debugger, Chrome DevTools, or Visual Studio Code`,
				Args: []model.Arg{{
					Name:       "port",
					IsOptional: true,
				}},
			}, {
				Name:        []string{"--inspect-brk"},
				Description: `Enable server-side debugging via debugging clients like the Node.js command-line debugger, Chrome DevTools, or Visual Studio Code. The server will be paused at startup, waiting for clients to attach to the process on the specified port`,
				Args: []model.Arg{{
					Name:       "port",
					IsOptional: true,
				}},
			}, {
				Name:        []string{"--mobile-server"},
				Description: `Location where mobile builds connect to the Meteor server. Defaults to your local IP and the port that the Meteor server binds to. Can include a URL scheme (for example, --mobile-server https://example.com:443)`,
			}, {
				Name:        []string{"--cordova-server-port"},
				Description: `Local port where Cordova will serve the content. It's important when multiple Cordova apps are build from the same Meteor app source code as by default the port is generated using the id inside .meteor/.id file`,
				Args: []model.Arg{{
					Name: "port",
				}},
			}, {
				Name:        []string{"--production"},
				Description: `Simulate production mode. Minify and bundle CSS and JS files`,
			}, {
				Name:        []string{"--raw-logs"},
				Description: `Run without parsing logs from stdout and stderr`,
			}, {
				Name:        []string{"--settings", "-s"},
				Description: `Set optional data for Meteor.settings on the server`,
				Args: []model.Arg{{
					Name:      "JSON File",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--release"},
				Description: `Specify the release of Meteor to use`,
				Args: []model.Arg{{
					Name: "release",
				}},
			}, {
				Name:        []string{"--verbose"},
				Description: `Print all output from builds logs`,
			}, {
				Name:        []string{"--no-lint"},
				Description: `Don't run linters used by the app on every rebuild`,
			}, {
				Name:        []string{"--no-release-check"},
				Description: `Don't run the release updater to check for new releases`,
			}, {
				Name:        []string{"--allow-incompatible-update"},
				Description: `Allow packages in your project to be upgraded or downgraded to versions that are potentially incompatible with the current versions, if required to satisfy all package version constraints`,
			}, {
				Name:        []string{"--extra-packages"},
				Description: `Run with additional packages (comma separated, for example: --extra-packages "package-name1, package-name2@1.2.3")`,
				Args: []model.Arg{{
					Name: "packages",
				}},
			}, {
				Name:        []string{"--exclude-archs"},
				Description: `Don't create bundles for certain web architectures (comma separated, for example: --exclude-archs "web.browser.legacy, web.cordova")`,
				Args: []model.Arg{{
					Name: "architectures",
				}},
			}},
		}, {
			Name:        []string{"debug"},
			Description: `DEPRECATED. The 'meteor debug' command is deprecated in favor of 'meteor --inspect-brk'`,
		}, {
			Name:        []string{"create"},
			Description: `Create a new Meteor project. By default, it uses React and makes a subdirectory named name and copies in the template app. You can pass an absolute or relative path`,
			Args: []model.Arg{{
				Name: "path",
			}},
			Options: []model.Option{{
				Name:        []string{"--release"},
				Description: `Specify the release of Meteor to use`,
				Args: []model.Arg{{
					Name: "release",
				}},
			}, {
				Name:        []string{"--package"},
				Description: `Create a new meteor package instead of an app`,
				Args: []model.Arg{{
					Name: "package_name",
				}},
			}, {
				Name:        []string{"--example"},
				Description: `Example template to use`,
				Args: []model.Arg{{
					Name:      "example_name",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--list"},
				Description: `Show list of available examples`,
			}, {
				Name:        []string{"--bare"},
				Description: `Create an empty app`,
			}, {
				Name:        []string{"--minimal"},
				Description: `Create an app with as few Meteor packages as possible`,
			}, {
				Name:        []string{"--full"},
				Description: `Create a fully scaffolded app`,
			}, {
				Name:        []string{"--react"},
				Description: `Create a basic react-based app, same as default`,
			}, {
				Name:        []string{"--vue"},
				Description: `Create a basic vue-based app`,
			}, {
				Name:        []string{"--apollo"},
				Description: `Create a basic apollo-based app`,
			}, {
				Name:        []string{"--svelte"},
				Description: `Create a basic svelte-based app`,
			}, {
				Name:        []string{"--typescript"},
				Description: `Create a basic Typescript React-based app`,
			}, {
				Name:        []string{"--blaze"},
				Description: `Create a basic blaze-based app`,
			}, {
				Name:        []string{"--tailwind"},
				Description: `Create a basic react-based app, with tailwind configured`,
			}},
		}, {
			Name:        []string{"update"},
			Description: `Attempts to bring you to the latest version of Meteor, and then to upgrade your packages to their latest versions. By default, update will not break compatibility`,
			Args: []model.Arg{{
				Name:       "package_names",
				Generator:  nil, // TODO: port over generator
				IsOptional: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--packages-only"},
				Description: `Update the package versions only. Do not update the release`,
			}, {
				Name:        []string{"--patch"},
				Description: `Update the release to a patch release only`,
			}, {
				Name:        []string{"--release"},
				Description: `Update to a specific release of meteor`,
				Args: []model.Arg{{
					Name: "release",
				}},
			}, {
				Name:        []string{"--allow-incompatible-update"},
				Description: `Allow packages in your project to be upgraded or downgraded to versions that are potentially incompatible with the current versions, if required to satisfy all package version constraints`,
			}, {
				Name:        []string{"--all-packages"},
				Description: `Update all packages including indirect dependencies`,
			}},
		}, {
			Name:        []string{"add"},
			Description: `Add packages to your Meteor project. You can add multiple packages with one command`,
			Args: []model.Arg{{
				Name: "package",
			}},
			Options: []model.Option{{
				Name:        []string{"--allow-incompatible-update"},
				Description: `Allow packages in your project to be upgraded or downgraded to versions that are potentially incompatible with the current versions, if required to satisfy all package version constraints`,
			}},
		}, {
			Name:        []string{"remove"},
			Description: `Removes a package previously added to your Meteor project. You can remove multiple packages with one command. For a list of the packages that your application is currently using, see "meteor list"`,
			Args: []model.Arg{{
				Name:      "package",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"list"},
			Description: `Lists the packages that you have explicitly added to your project`,
			Options: []model.Option{{
				Name:        []string{"--tree"},
				Description: `Outputs a tree showing how packages are referenced`,
			}, {
				Name:        []string{"--json"},
				Description: `Outputs the tree in JSON format`,
			}, {
				Name:        []string{"--weak"},
				Description: `Show weakly referenced dependencies in the tree`,
			}, {
				Name:        []string{"--details"},
				Description: `Adds more package details to the JSON output`,
			}},
		}, {
			Name:        []string{"add-platform"},
			Description: `Adds platforms to your Meteor project. You can add multiple platforms with one command`,
			Args: []model.Arg{{
				Name:        "platform",
				Suggestions: []model.Suggestion{{Name: []string{`server`}}, {Name: []string{`browser`}}, {Name: []string{`android`}}, {Name: []string{`ios`}}},
			}},
		}, {
			Name:        []string{"remove-platform"},
			Description: `Removes a platform previously added to your Meteor project`,
			Args: []model.Arg{{
				Name:      "platform",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"list-platforms"},
			Description: `Lists all of the platforms added to your project`,
		}, {
			Name:        []string{"ensure-cordova-dependencies"},
			Description: `Check if the dependencies are installed, otherwise install them`,
		}, {
			Name:        []string{"build"},
			Description: `Package this project up for deployment. The command outputs a directory with builds for all platforms in this project`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "output path",
			}},
			Options: []model.Option{{
				Name:        []string{"--debug"},
				Description: `Build in debug mode (don't minify, etc)`,
			}, {
				Name:        []string{"--directory"},
				Description: `Output a directory (rather than a tarball) for the application server bundle. If the output location exists, it will be recursively deleted first`,
			}, {
				Name:        []string{"--server-only"},
				Description: `Skip building mobile apps even if mobile platforms have been added`,
			}, {
				Name:        []string{"--mobile-settings"},
				Description: `Set optional data for the initial value of Meteor.settings in your mobile application`,
			}, {
				Name:        []string{"--server"},
				Description: `Location where mobile builds connect to the Meteor server. Defaults to localhost:3000. Can include a URL scheme (for example, --server https://example.com:443)`,
			}, {
				Name:        []string{"--architecture"},
				Description: `Builds the server for a different architecture than your developer machine's architecture. Note: This option selects the architecture of the binary-dependent Atmosphere packages you would like bundled into your application, when those packages were specifically published for multiple architectures (i.e. with meteor publish-for-arch). If your project doesn't use any Atmosphere packages that have binary dependencies, --architecture has no effect`,
				Args: []model.Arg{{
					Name:        "arch",
					Suggestions: []model.Suggestion{{Name: []string{`os.osx.x86_64`}}, {Name: []string{`os.linux.x86_64`}}, {Name: []string{`os.linux.x86_32`}}, {Name: []string{`os.windows.x86_32`}}, {Name: []string{`os.windows.x86_64`}}},
				}},
			}, {
				Name:        []string{"--allow-incompatible-update"},
				Description: `Allow packages in your project to be upgraded or downgraded to versions that are potentially incompatible with the current versions, if required to satisfy all package version constraints`,
			}, {
				Name:        []string{"--platforms"},
				Description: `Builds only the specified platforms (when available)`,
				Args: []model.Arg{{
					Name:       "platforms",
					Generator:  nil, // TODO: port over generator
					IsOptional: true,
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--packageType"},
				Description: `Choose between apk/bundle for android builds`,
				Args: []model.Arg{{
					Name:        "package-type",
					Suggestions: []model.Suggestion{{Name: []string{`apk`}}, {Name: []string{`bundle`}}},
					IsOptional:  true,
				}},
			}},
		}, {
			Name:        []string{"lint"},
			Description: `Run through the whole build process for the app and run all linters the app uses`,
			Options: []model.Option{{
				Name:        []string{"--allow-incompatible-update"},
				Description: `Allow packages in your project to be upgraded or downgraded to versions that are potentially incompatible with the current versions, if required to satisfy all package version constraints`,
			}},
		}, {
			Name:        []string{"shell"},
			Description: `When "meteor shell" is executed in an application directory where a server is already running, it connects to the server and starts an interactive shell for evaluating server-side code`,
		}, {
			Name:        []string{"mongo"},
			Description: `Opens a Mongo shell to view or manipulate collections. Can only be used when developing locally`,
			Options: []model.Option{{
				Name:        []string{"--url", "-U"},
				Description: `Mongo database URL`,
				Args: []model.Arg{{
					Name: "url",
				}},
			}},
		}, {
			Name:        []string{"reset"},
			Description: `Reset the current project to a fresh state. Removes all local data`,
		}, {
			Name:        []string{"deploy"},
			Description: `Deploys the project in your current directory to Meteor's servers`,
			Args: []model.Arg{{
				Name:        "site",
				Description: `You can deploy to any available name under "meteorapp.com" without any additional configuration, for example, "myapp.meteorapp.com". If you deploy to a custom domain, such as "myapp.mydomain.com", then you'll also need to configure your domain's DNS records. See the Meteor / Galaxy docs (http://cloud-guide.meteor.com/dns.html) for details`,
			}},
			Options: []model.Option{{
				Name:        []string{"--delete", "-D"},
				Description: `Permanently delete this deployment`,
			}, {
				Name:        []string{"--debug"},
				Description: `Deploy in debug mode (don't minify, etc)`,
			}, {
				Name:        []string{"--settings", "-s"},
				Description: `Set optional data for Meteor.settings`,
				Args: []model.Arg{{
					Name:      "JSON File",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--allow-incompatible-update"},
				Description: `Allow packages in your project to be upgraded or downgraded to versions that are potentially incompatible with the current versions, if required to satisfy all package version constraints`,
			}, {
				Name:        []string{"--deploy-polling-timeout"},
				Description: `The number of milliseconds to wait for build/deploy success or failure after a successful upload of your app's minified code; defaults to 15 minutes`,
				Args: []model.Arg{{
					Name:        "milliseconds",
					Description: `Milliseconds to wait`,
				}},
			}, {
				Name:        []string{"--no-wait"},
				Description: `Exits when Meteor has uploaded the app's code instead of waiting for the deploy to conclude`,
			}, {
				Name:        []string{"--cache-build"},
				Description: `Reuses the build already created if the git commit hash is the same`,
			}, {
				Name:        []string{"--free"},
				Description: `When deploying an app for the first time, you can pass this option to deploy your app in the Galaxy's free mode`,
			}, {
				Name:        []string{"--plan"},
				Description: `You can change the app plan by providing this argument with one of the following values: professional, essentials, or free. Be aware that this argument overwrites the "--free" argument`,
			}, {
				Name:        []string{"--mongo"},
				Description: `If this flag is true and it's not provided a mongo url in the settings ('galaxy.meteor.com'.env.MONGO_URL), when deploying, Galaxy will create a database to your app in its shared cluster and will insert the URL in your app's settings for you`,
				Args: []model.Arg{{
					Name:        "true/false",
					Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}},
				}},
			}, {
				Name:        []string{"--container-size"},
				Description: `Change your app's container size using the deploy command`,
				Args: []model.Arg{{
					Name:        "size",
					Suggestions: []model.Suggestion{{Name: []string{`tiny`}}, {Name: []string{`compact`}}, {Name: []string{`standard`}}, {Name: []string{`double`}}, {Name: []string{`quad`}}, {Name: []string{`octa`}}, {Name: []string{`dozen`}}},
				}},
			}},
		}, {
			Name:        []string{"authorized"},
			Description: `Manage authorized users and organizations`,
			Args: []model.Arg{{
				Name: "site",
			}},
			Options: []model.Option{{
				Name:        []string{"--add"},
				Description: `Add an authorized user or organization`,
			}, {
				Name:        []string{"--remove"},
				Description: `Remove an authorized user or organization`,
			}, {
				Name:        []string{"--transfer"},
				Description: `Transfer the (Galaxy) app to a new user or organization`,
			}, {
				Name:        []string{"--list"},
				Description: `List authorized users and organizations`,
			}},
		}, {
			Name:        []string{"login"},
			Description: `Prompts for your username and password and logs you in to your Meteor developer account`,
			Options: []model.Option{{
				Name:        []string{"--email"},
				Description: `Add an authorized user or organization`,
				Args: []model.Arg{{
					Name: "email",
				}},
			}},
		}, {
			Name:        []string{"logout"},
			Description: `Log out of your Meteor developer account`,
		}, {
			Name:        []string{"whoami"},
			Description: `Prints the username of the currently logged-in Meteor developer`,
		}, {
			Name:        []string{"test-packages"},
			Description: `Runs unit tests for one or more packages. The results are shown in a browser dashboard that updates whenever a relevant source file is modified`,
			Options: []model.Option{{
				Name:        []string{"--release"},
				Description: `Specify the release of Meteor to use`,
				Args: []model.Arg{{
					Name: "release",
				}},
			}, {
				Name:        []string{"--port", "-p"},
				Description: `Port to listen on (instead of the default 3000). Also uses port N+1 and N+2`,
				Args: []model.Arg{{
					Name: "port",
				}},
			}, {
				Name:        []string{"--inspect"},
				Description: `Enable server-side debugging via debugging clients like the Node.js command-line debugger, Chrome DevTools, or Visual Studio Code`,
				Args: []model.Arg{{
					Name:       "port",
					IsOptional: true,
				}},
			}, {
				Name:        []string{"--inspect-brk"},
				Description: `Enable server-side debugging via debugging clients like the Node.js command-line debugger, Chrome DevTools, or Visual Studio Code. The server will be paused at startup, waiting for clients to attach to the process on the specified port`,
				Args: []model.Arg{{
					Name:       "port",
					IsOptional: true,
				}},
			}, {
				Name:        []string{"--mobile-server"},
				Description: `Location where mobile builds connect to the Meteor server. Defaults to your local IP and the port that the Meteor server binds to. Can include a URL scheme (for example, --mobile-server https://example.com:443)`,
			}, {
				Name:        []string{"--cordova-server-port"},
				Description: `Local port where Cordova will serve the content. It's important when multiple Cordova apps are build from the same Meteor app source code as by default the port is generated using the id inside .meteor/.id file`,
				Args: []model.Arg{{
					Name: "port",
				}},
			}, {
				Name:        []string{"--production"},
				Description: `Simulate production mode. Minify and bundle CSS and JS files`,
			}, {
				Name:        []string{"--settings", "-s"},
				Description: `Set optional data for Meteor.settings on the server`,
				Args: []model.Arg{{
					Name:      "JSON File",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--ios"},
				Description: `Run tests in an iOS emulator. All of the tests for client and server will run in addition to mobile-specific tests`,
			}, {
				Name:        []string{"--android"},
				Description: `Run tests in an Android emulator. All of the tests for client and server will run in addition to mobile-specific tests`,
			}, {
				Name:        []string{"--ios-device"},
				Description: `Run tests on iOS mobile device. All of the tests for client and server will run in addition to mobile-specific tests`,
			}, {
				Name:        []string{"--android-device"},
				Description: `Run tests on Android emulator or on a mobile device. All of the tests for client and server will run in addition to mobile-specific tests`,
			}, {
				Name:        []string{"--test-app-path"},
				Description: `Set the directory in which to create a temporary app used for tests. Defaults to the system's temporary directory, usually /tmp`,
				Args: []model.Arg{{
					Templates:  []model.Template{model.TemplateFolders},
					Name:       "path",
					IsOptional: true,
				}},
			}, {
				Name:        []string{"--verbose"},
				Description: `Print all output from builds logs`,
			}, {
				Name:        []string{"--no-lint"},
				Description: `Don't run linters used by the app on every rebuild`,
			}, {
				Name:        []string{"--extra-packages"},
				Description: `Run with additional packages (comma separated, for example: --extra-packages "package-name1, package-name2@1.2.3")`,
				Args: []model.Arg{{
					Name: "packages",
				}},
			}, {
				Name:        []string{"--driver-package"},
				Description: `Test driver package to use to run tests and display results. For example: --driver-package meteortesting:mocha`,
				Args: []model.Arg{{
					Name: "driver-package",
				}},
			}, {
				Name:        []string{"package"},
				Description: `Package name`,
			}},
		}, {
			Name:        []string{"test"},
			Description: `Runs tests against the application`,
			Options: []model.Option{{
				Name:        []string{"--full-app"},
				Description: `Your app is loaded as usual, then hidden, so that your tests can inspect and interact with the full running application`,
				Args: []model.Arg{{
					Name: "port",
				}},
			}, {
				Name:        []string{"--port", "-p"},
				Description: `Port to listen on (instead of the default 3000). Also uses port N+1 and N+2`,
				Args: []model.Arg{{
					Name: "port",
				}},
			}, {
				Name:        []string{"--inspect"},
				Description: `Enable server-side debugging via debugging clients like the Node.js command-line debugger, Chrome DevTools, or Visual Studio Code`,
				Args: []model.Arg{{
					Name:       "port",
					IsOptional: true,
				}},
			}, {
				Name:        []string{"--inspect-brk"},
				Description: `Enable server-side debugging via debugging clients like the Node.js command-line debugger, Chrome DevTools, or Visual Studio Code. The server will be paused at startup, waiting for clients to attach to the process on the specified port`,
				Args: []model.Arg{{
					Name:       "port",
					IsOptional: true,
				}},
			}, {
				Name:        []string{"--mobile-server"},
				Description: `Location where mobile builds connect to the Meteor server. Defaults to your local IP and the port that the Meteor server binds to. Can include a URL scheme (for example, --mobile-server https://example.com:443)`,
			}, {
				Name:        []string{"--cordova-server-port"},
				Description: `Local port where Cordova will serve the content. It's important when multiple Cordova apps are build from the same Meteor app source code as by default the port is generated using the id inside .meteor/.id file`,
				Args: []model.Arg{{
					Name: "port",
				}},
			}, {
				Name:        []string{"--raw-logs"},
				Description: `Run without parsing logs from stdout and stderr`,
			}, {
				Name:        []string{"--settings", "-s"},
				Description: `Set optional data for Meteor.settings on the server`,
				Args: []model.Arg{{
					Name:      "JSON File",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--ios"},
				Description: `Run tests in an iOS emulator. All of the tests for client and server will run in addition to mobile-specific tests`,
			}, {
				Name:        []string{"--android"},
				Description: `Run tests in an Android emulator. All of the tests for client and server will run in addition to mobile-specific tests`,
			}, {
				Name:        []string{"--ios-device"},
				Description: `Run tests on iOS mobile device. All of the tests for client and server will run in addition to mobile-specific tests`,
			}, {
				Name:        []string{"--android-device"},
				Description: `Run tests on Android emulator or on a mobile device. All of the tests for client and server will run in addition to mobile-specific tests`,
			}, {
				Name:        []string{"--test-app-path"},
				Description: `Set the directory in which to create a temporary app used for tests. Defaults to the system's temporary directory, usually /tmp`,
				Args: []model.Arg{{
					Templates:  []model.Template{model.TemplateFolders},
					Name:       "path",
					IsOptional: true,
				}},
			}, {
				Name:        []string{"--verbose"},
				Description: `Print all output from builds logs`,
			}, {
				Name:        []string{"--extra-packages"},
				Description: `Run with additional packages (comma separated, for example: --extra-packages "package-name1, package-name2@1.2.3")`,
				Args: []model.Arg{{
					Name: "packages",
				}},
			}, {
				Name:        []string{"--driver-package"},
				Description: `Test driver package to use to run tests and display results. For example: --driver-package meteortesting:mocha`,
				Args: []model.Arg{{
					Name: "driver-package",
				}},
			}, {
				Name:        []string{"--exclude-archs"},
				Description: `Don't create bundles for certain web architectures (comma separated, for example: --exclude-archs "web.browser.legacy, web.cordova")`,
				Args: []model.Arg{{
					Name: "architectures",
				}},
			}},
		}, {
			Name:        []string{"admin"},
			Description: `Rarely used commands for administering official Meteor services`,
			Args: []model.Arg{{
				Name:      "command",
				IsCommand: true,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"maintainers"},
				Description: `Manage users and organizations that are maintainers for a particular package`,
				Args: []model.Arg{{
					Name: "package_name",
				}},
				Options: []model.Option{{
					Name:        []string{"--add"},
					Description: `Add an authorized maintainer to a package`,
					Args: []model.Arg{{
						Name: "username",
					}},
				}, {
					Name:        []string{"--remove"},
					Description: `Remove an authorized maintainer from a package. You cannot remove yourself if you are the last maintainer on a package`,
					Args: []model.Arg{{
						Name: "username",
					}},
				}, {
					Name:        []string{"--list"},
					Description: `List authorized maintainers`,
				}},
			}, {
				Name:        []string{"recommend-release"},
				Description: `Recommend a version of a meteor release`,
				Args: []model.Arg{{
					Name: "package_name",
				}},
			}, {
				Name:        []string{"change-homepage"},
				Description: `Change the homepage containing package information`,
				Args: []model.Arg{{
					Name: "package_name",
				}, {
					Name: "new_url",
				}},
			}, {
				Name:        []string{"list-organizations"},
				Description: `List the organizations of which you are a member`,
			}, {
				Name:        []string{"members"},
				Description: `Manage members of an organization`,
				Args: []model.Arg{{
					Name: "organization_name",
				}},
				Options: []model.Option{{
					Name:        []string{"--add"},
					Description: `Add a member to the organization`,
					Args: []model.Arg{{
						Name: "username",
					}},
				}, {
					Name:        []string{"--remove"},
					Description: `Remove a member to the organization`,
					Args: []model.Arg{{
						Name: "username",
					}},
				}, {
					Name:        []string{"--list"},
					Description: `List members of the organization`,
				}},
			}},
		}, {
			Name:        []string{"list-sites"},
			Description: `List the sites that you have deployed with "meteor deploy", and sites for which other users have authorized you with the "meteor authorized" command. To see sites in a region other than us-east-1, set the DEPLOY_HOSTNAME environment variable. For example, "DEPLOY_HOSTNAME=eu-west-1.galaxy-deploy.meteor.com meteor list-sites"`,
		}, {
			Name:        []string{"publish-release"},
			Description: `Publishes a new release to the package server, as determined by a JSON configuration file`,
			Args: []model.Arg{{
				Name:      "JSON File",
				Generator: nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"--create-track"},
				Description: `Publish a new release track`,
			}},
		}, {
			Name:        []string{"publish"},
			Description: `Publishes a new version of a local package to the package server. Must be runfrom the directory containing the package`,
			Options: []model.Option{{
				Name:        []string{"--create"},
				Description: `Publish a new package`,
			}, {
				Name:        []string{"--update"},
				Description: `Publish a new package`,
			}, {
				Name:        []string{"--allow-incompatible-update"},
				Description: `Allow packages in your project to be upgraded or downgraded to versions that are potentially incompatible with the current versions, if required to satisfy all package version constraints`,
			}, {
				Name:        []string{"--no-lint"},
				Description: `Don't run linters used by the app on every rebuild`,
			}},
		}, {
			Name:        []string{"publish-for-arch"},
			Description: `Publish a package for the same architecture that you are currently on`,
			Args: []model.Arg{{
				Name: "packageName@version",
			}},
		}, {
			Name:        []string{"search"},
			Description: `Searches through the Meteor package and release database for items whose names match the regular expression`,
			Args: []model.Arg{{
				Name: "regex",
			}},
			Options: []model.Option{{
				Name:        []string{"--maintainer"},
				Description: `Filter by authorized maintainer`,
				Args: []model.Arg{{
					Name: "maintainer_name",
				}},
			}, {
				Name:        []string{"--show-all"},
				Description: `Show all matches, even prereleases`,
			}, {
				Name:        []string{"--ejson"},
				Description: `Show more detailed output in EJSON format`,
			}},
		}, {
			Name:        []string{"show"},
			Description: `Searches through the Meteor package and release database for items whose names match the regular expression`,
			Args: []model.Arg{{
				Name: "name or name@version",
			}},
			Options: []model.Option{{
				Name:        []string{"--show-all"},
				Description: `Show all matches, even prereleases`,
			}, {
				Name:        []string{"--ejson"},
				Description: `Show more detailed output in EJSON format`,
			}},
		}},
	}
}
