package cmd

import (
	"github.com/spf13/cobra"

	"github.com/craftcms/nitro/config"
	"github.com/craftcms/nitro/internal/action"
	"github.com/craftcms/nitro/internal/nitro"
)

var sshCommand = &cobra.Command{
	Use:   "ssh",
	Short: "SSH into machine",
	RunE: func(cmd *cobra.Command, args []string) error {
		name := config.GetString("name", flagMachineName)

		sshAction, err := action.SSH(name)
		if err != nil {
			return err
		}

		return nitro.RunAction(nitro.NewMultipassRunner("multipass"), []action.Action{*sshAction})
	},
}
