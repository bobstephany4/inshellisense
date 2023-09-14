// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["kind"] = model.Subcommand{
		Name:        []string{"kind"},
		Description: `Kubernetes IN Docker - local clusters for testing Kubernetes`,
		Options: []model.Option{{
			Name:         []string{"-h", "--help"},
			Description:  `Help for kind`,
			IsPersistent: true,
		}, {
			Name:         []string{"-q", "--quiet"},
			Description:  `Silence all stderr output`,
			IsPersistent: true,
		}, {
			Name:        []string{"-v", "--verbosity"},
			Description: `Info log verbosity, higher value produces more output`,
			Args: []model.Arg{{
				Name: "int",
			}},
			IsPersistent: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"Build"},
			Description: `Build one of [node-image]`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"node-image"},
				Description: `Builds a node image`,
				Options: []model.Option{{
					Name:        []string{"--arch"},
					Description: `Architecture to build for, defaults to the host architecture`,
					Args: []model.Arg{{
						Name:        "architecture",
						Description: `Architecture to build for, defaults to the host architecture`,
					}},
				}, {
					Name:        []string{"--base-image"},
					Description: `Base image to use for the node image`,
					Args: []model.Arg{{
						Name:        "base image",
						Description: `Name:tag of the base image to use for the build`,
					}},
				}, {
					Name:        []string{"--image"},
					Description: `Name:tag of the resulting image to be built`,
					Args: []model.Arg{{
						Name:        "name:tag",
						Description: `Name:tag of the resulting image to be built`,
					}},
				}, {
					Name:        []string{"--kube-root"},
					Description: `Path to the Kubernetes source directory`,
					Args: []model.Arg{{
						Name:        "path",
						Description: `Path to the Kubernetes source directory`,
					}},
				}, {
					Name:        []string{"--type"},
					Description: `Type of node image to build`,
					Args: []model.Arg{{
						Name:        "type",
						Description: `Build type, default is docker`,
					}},
				}},
			}},
		}, {
			Name:        []string{"create"},
			Description: `Creates a cluster`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"cluster"},
				Description: `Creates a cluster`,
				Options: []model.Option{{
					Name:        []string{"--config"},
					Description: `Path to a kind config file`,
				}, {
					Name:        []string{"--image"},
					Description: `Node docker image to use for booting the cluster`,
				}, {
					Name:        []string{"--kubeconfig"},
					Description: `Sets kubeconfig path instead of $KUBECONFIG or $HOME/.kube/config`,
				}, {
					Name:        []string{"--name"},
					Description: `Cluster name`,
				}, {
					Name:        []string{"--retain"},
					Description: `Retain nodes for debugging when cluster creation fails`,
				}, {
					Name:        []string{"--wait"},
					Description: `Wait for control-plane node to be ready`,
				}},
			}},
		}, {
			Name:        []string{"completion"},
			Description: `Generates shell completion scripts`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"bash"},
				Description: `Output shell completions for bash`,
			}, {
				Name:        []string{"fish"},
				Description: `Output shell completions for fish`,
			}, {
				Name:        []string{"zsh"},
				Description: `Output shell completions for zsh`,
			}},
		}, {
			Name:        []string{"delete"},
			Description: `Deletes one or more clusters`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"cluster"},
				Description: `Delete Cluster`,
				Options: []model.Option{{
					Name:        []string{"--name"},
					Description: `Cluster name`,
					Args: []model.Arg{{
						Name:      "cluster name",
						Generator: nil, // TODO: port over generator
					}},
				}},
			}, {
				Name:        []string{"clusters"},
				Description: `Delete Clusters`,
				Options: []model.Option{{
					Name:        []string{"-A", "--all"},
					Description: `Delete all clusters`,
				}, {
					Name:        []string{"--kubeconfig"},
					Description: `Sets kubeconfig path instead of $KUBECONFIG or $HOME/.kube/config`,
				}},
			}},
		}, {
			Name:        []string{"export"},
			Description: `Exports a cluster's kubeconfig`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"kubeconfig"},
				Description: `Exports a cluster's kubeconfig`,
				Options: []model.Option{{
					Name:        []string{"--name"},
					Description: `Cluster name`,
				}, {
					Name:        []string{"--internal"},
					Description: `Use internal address instead of externalt`,
				}, {
					Name:        []string{"--kubeconfig"},
					Description: `Sets kubeconfig path instead of $KUBECONFIG or $HOME/.kube/config`,
				}},
			}, {
				Name:        []string{"logs"},
				Description: `Exports logs to a tempdir or [output-dir] if specified`,
				Options: []model.Option{{
					Name:        []string{"--name"},
					Description: `Cluster name`,
					Args: []model.Arg{{
						Name:      "cluster name",
						Generator: nil, // TODO: port over generator
					}},
				}},
			}},
		}, {
			Name:        []string{"get"},
			Description: `Gets one of [clusters, nodes, kubeconfig]`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"clusters"},
				Description: `Lists existing kind clusters by their name`,
			}, {
				Name:        []string{"kubeconfig"},
				Description: `Prints cluster kubeconfig`,
				Options: []model.Option{{
					Name:        []string{"--name"},
					Description: `Cluster name`,
					Args: []model.Arg{{
						Name:      "cluster name",
						Generator: nil, // TODO: port over generator
					}},
				}, {
					Name:        []string{"--internal"},
					Description: `Use internal address instead of external`,
				}},
			}, {
				Name:        []string{"nodes"},
				Description: `Lists existing kind nodes by their name`,
				Options: []model.Option{{
					Name:        []string{"-A", "--all"},
					Description: `List all nodes`,
				}, {
					Name:        []string{"--name"},
					Description: `Cluster name`,
					Args: []model.Arg{{
						Name:      "cluster name",
						Generator: nil, // TODO: port over generator
					}},
				}},
			}},
		}, {
			Name:        []string{"load"},
			Description: `Loads images into node from an archive or image on host`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"docker-image"},
				Description: `Loads docker images from host into all or specified nodes by name`,
				Options: []model.Option{{
					Name:        []string{"--name"},
					Description: `Cluster name`,
					Args: []model.Arg{{
						Name:      "cluster name",
						Generator: nil, // TODO: port over generator
					}},
				}, {
					Name:        []string{"--nodes"},
					Description: `Comma separated list of nodes to load images into`,
					Args: []model.Arg{{
						Name:      "nodes",
						Generator: nil, // TODO: port over generator
					}},
				}},
			}, {
				Name:        []string{"image-archive"},
				Description: `Loads docker images from archive into all or specified nodes by name`,
				Options: []model.Option{{
					Name:        []string{"--name"},
					Description: `Cluster name`,
					Args: []model.Arg{{
						Name:      "cluster name",
						Generator: nil, // TODO: port over generator
					}},
				}, {
					Name:        []string{"--nodes"},
					Description: `Comma separated list of nodes to load images into`,
					Args: []model.Arg{{
						Name:      "nodes",
						Generator: nil, // TODO: port over generator
					}},
				}},
			}},
		}, {
			Name:        []string{"version"},
			Description: `Prints the kind CLI version`,
		}},
	}
}
