package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommand(name string) *cobra.Command {

	c := &cobra.Command{
		Use:   fmt.Sprint("%s", name),
		Short: "Program used to execute transaction on the Ethereum Blockchain when certain crietria are met",
		Long:  "Program used to execute transaction on the Ethereum Blockchain when certain crietria are met",
	}

	c.AddCommand(
		NewCreateCommand(),
	)

	return c
}
