/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package asciidoc

import (
	"github.com/spf13/cobra"

	"github.com/terraform-docs/terraform-docs/cmd/asciidoc/document"
	"github.com/terraform-docs/terraform-docs/cmd/asciidoc/table"
	"github.com/terraform-docs/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for 'asciidoc' formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "asciidoc [PATH]",
		Aliases:     []string{"adoc"},
		Short:       "Generate AsciiDoc of inputs and outputs",
		Annotations: cli.Annotations("asciidoc"),
		PreRunE:     cli.PreRunEFunc(config),
		RunE:        cli.RunEFunc(config),
	}

	// flags
	cmd.PersistentFlags().BoolVar(&config.Settings.Required, "required", true, "show Required column or section")
	cmd.PersistentFlags().BoolVar(&config.Settings.Sensitive, "sensitive", true, "show Sensitive column or section")
	cmd.PersistentFlags().IntVar(&config.Settings.Indent, "indent", 2, "indention level of AsciiDoc sections [1, 2, 3, 4, 5]")
	cmd.PersistentFlags().BoolVar(&config.Settings.Anchor, "anchor", true, "create anchor links")

	// subcommands
	cmd.AddCommand(document.NewCommand(config))
	cmd.AddCommand(table.NewCommand(config))

	return cmd
}
