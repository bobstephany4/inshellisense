// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["cdk"] = model.Subcommand{
		Name:        []string{"cdk"},
		Description: `AWS CDK CLI`,
		Options: []model.Option{{
			Name:        []string{"--version"},
			Description: `The current version`,
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Show help`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"init"},
			Description: `Create a new, empty CDK project from a template`,
		}, {
			Name:        []string{"metadata"},
			Description: `Returns all metadata associated with this stack`,
		}, {
			Name:        []string{"doctor"},
			Description: `Check your set-up for potential problems`,
		}, {
			Name:        []string{"diff"},
			Description: `Compares the specified stack with the deployed stack`,
		}, {
			Name:        []string{"destroy"},
			Description: `Destroy the specified stack(s)`,
		}, {
			Name:        []string{"deploy"},
			Description: `Deploy the specified stack(s) into your AWS account`,
		}, {
			Name:        []string{"bootstrap"},
			Description: `Deploys the CDK toolkit stack into an AWS environment`,
		}, {
			Name:        []string{"synth", "synthesize"},
			Description: `Synthesizes and prints the CloudFormation template for this stack`,
		}, {
			Name:        []string{"ls", "list"},
			Description: `List all stacks in the app`,
		}, {
			Name:        []string{"import"},
			Description: `Import existing resource(s) into the given STACK`,
		}, {
			Name:        []string{"watch"},
			Description: `Shortcut for 'deploy --watch'`,
		}, {
			Name:        []string{"ack", "acknowledge"},
			Description: `Acknowledge a notice so that it does not show up anymore`,
		}, {
			Name:        []string{"notices"},
			Description: `Returns a list of relevant notices`,
		}, {
			Name:        []string{"context"},
			Description: `Manage cached context values`,
		}, {
			Name:        []string{"doc", "docs"},
			Description: `Opens the reference documentation in a browser`,
		}},
	}
}
