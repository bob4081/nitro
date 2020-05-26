package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/craftcms/nitro/internal/nitro"
	"github.com/craftcms/nitro/internal/scripts"
)

var mailhogCommand = &cobra.Command{
	Use:   "mailhog",
	Short: "Install mailhog",
	RunE: func(cmd *cobra.Command, args []string) error {
		machine := flagMachineName
		mp, err := exec.LookPath("multipass")
		if err != nil {
			return err
		}
		script := scripts.New(mp, machine)
		ip := nitro.IP(machine, nitro.NewMultipassRunner(mp))

		// check if the container is already running
		if output, err := script.Run(false, fmt.Sprintf(scripts.FmtDockerContainerExists, "mailhog")); err == nil {
			if strings.Contains(output, "exists") {
				fmt.Println(fmt.Sprintf("Mailhog is already running, try accessing http://%s:8025", ip))
				return nil
			}
		}

		// create the local directory for mailhog
		if output, err := script.Run(false, fmt.Sprintf(scripts.FmtCreateDirectory, "/home/ubuntu/.nitro/mailhog")); err != nil {
			fmt.Println(output)
			return err
		}

		fmt.Println("Created maildir for mailhog at /home/ubuntu/.nitro/mailhog")

		// run mailhog container
		if output, err := script.Run(false, `docker run --name mailhog --restart always -d -e "MH_STORAGE=maildir" -v /home/ubuntu/.nitro/mailhog:/maildir -p 1025:1025 -p 8025:8025 mailhog/mailhog`); err != nil {
			fmt.Println(output)
			return err
		}

		fmt.Println(fmt.Sprintf("Mailhog is now running on SMTP %s:1025, you can view mailhog at http://%s:8025", ip, ip))

		return nil
	},
}
