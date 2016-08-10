package dev

import (
	"github.com/spf13/cobra"

	"github.com/nanobox-io/nanobox/processor"
	"github.com/nanobox-io/nanobox/util/print"
	"github.com/nanobox-io/nanobox/validate"
)

// RunCmd ...
var RunCmd = &cobra.Command{
	Use:    "run",
	Short:  "Opens an dev container and starts all the code commands init.",
	Long:   ``,
	PreRun: validate.Requires("provider", "provider_up", "dev_isup"),
	Run:    runFn,
}

// runFn ...
func runFn(ccmd *cobra.Command, args []string) {
	processor.DefaultControl.Env = "dev"

	// if given an argument they wanted to run a console into a container
	// if no arguement is provided they wanted to run a dev console
	// and be dropped into a dev environment
	processor.DefaultControl.Meta["run"] = "true"

	// set the meta arguments to be used in the processor and run the processor
	print.OutputCommandErr(processor.Run("dev_console", processor.DefaultControl))
}
