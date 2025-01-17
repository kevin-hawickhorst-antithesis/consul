// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package main

import (
	"fmt"
	"io"
	"log"
	"os"

	mcli "github.com/mitchellh/cli"

	"github.com/hashicorp/consul/command"
	"github.com/hashicorp/consul/command/cli"
	"github.com/hashicorp/consul/command/version"
	_ "github.com/hashicorp/consul/service_os"
	"github.com/TheBenCollins/antassert"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	log.SetOutput(io.Discard)

	ui := &cli.BasicUI{
		BasicUi: mcli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr},
	}
	cmds := command.RegisteredCommands(ui)

	// Assertion 1: Ensure that there are registered commands.
	antassert.AntAssert(len(cmds) > 0, "cmdRegistryNotEmpty", "No commands are registered.")


	var names []string
	for c := range cmds {
		names = append(names, c)
	}

	// Assertion 2: Ensure that the 'names' slice is not empty.
	antassert.AntAssert(len(names) > 0, "namesSliceNotEmpty", "The 'names' slice is empty.")


	cli := &mcli.CLI{
		Args:         os.Args[1:],
		Commands:     cmds,
		Autocomplete: true,
		Name:         "consul",
		HelpFunc:     mcli.FilteredHelpFunc(names, mcli.BasicHelpFunc("consul")),
		HelpWriter:   os.Stdout,
		ErrorWriter:  os.Stderr,
	}

	if cli.IsVersion() {
		cmd := version.New(ui)

		// Assertion 3: Ensure that the 'cmd' is not nil.
        antassert.AntAssert(cmd != nil, "versionCmdNotNil", "The 'version' command is nil.")

		return cmd.Run(nil)
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %v\n", err)
		return 1
	}

	// Assertion 4: Ensure that there is no error while executing the CLI.
    antassert.AntAssert(err == nil, "cliExecutionError", fmt.Sprintf("Error executing CLI: %v", err))

	return exitCode
}
