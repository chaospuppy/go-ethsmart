package main

import "github.com/spf13/cobra"

func NewCreateCommand() *cobra.Command {
	c := &cobra.Command{
		Use:       "create ",
		ValidArgs: []string{"transaction"},
		Short:     "Creates a resource",
		Args:      cobra.MaximumNArgs(1),
		Run: func(os.)
	}
}
