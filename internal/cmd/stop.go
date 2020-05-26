package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/craftcms/nitro/internal/nitro"
)

var stopCommand = &cobra.Command{
	Use:   "stop",
	Short: "Stop machine",
	RunE: func(cmd *cobra.Command, args []string) error {
		machine := flagMachineName

		stopAction, err := nitro.Stop(machine)
		if err != nil {
			return err
		}

		if err := nitro.Run(nitro.NewMultipassRunner("multipass"), []nitro.Action{*stopAction}); err != nil {
			return err
		}

		fmt.Println("Stopped", machine)

		return nil
	},
}
