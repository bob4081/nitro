package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/pixelandtonic/go-input"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/craftcms/nitro/config"
	"github.com/craftcms/nitro/internal/helpers"
	"github.com/craftcms/nitro/internal/nitro"
	"github.com/craftcms/nitro/internal/normalize"
	"github.com/craftcms/nitro/internal/prompt"
)

var importCommand = &cobra.Command{
	Use:   "import",
	Short: "Import database into a machine",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		machine := flagMachineName

		home, err := homedir.Dir()
		if err != nil {
			return err
		}

		// get the filename
		filename, fileAbsPath, err := normalize.Path(args[0], home)
		if err != nil {
			return err
		}

		// make sure the file exists
		if !helpers.FileExists(fileAbsPath) {
			return errors.New(fmt.Sprintf("Unable to locate the file %q", fileAbsPath))
		}

		// which database engine?
		var databases []config.Database
		if err := viper.UnmarshalKey("databases", &databases); err != nil {
			return err
		}
		var dbs []string
		for _, db := range databases {
			dbs = append(dbs, db.Name())
		}
		ui := &input.UI{
			Writer: os.Stdout,
			Reader: os.Stdin,
		}

		if len(dbs) == 0 {
			return errors.New("there are no databases that we can import the file into")
		}

		containerName, _, err := prompt.Select(ui, "Select a database engine to import the backup into", dbs[0], dbs)

		databaseName, err := prompt.Ask(ui, "What is the database name?", "", true)
		if err != nil {
			return err
		}

		var actions []nitro.Action

		// syntax is strange, see this issue: https://github.com/canonical/multipass/issues/1165#issuecomment-548763143
		transferAction := nitro.Action{
			Type:       "transfer",
			UseSyscall: false,
			Args:       []string{"transfer", fileAbsPath, machine + ":" + filename},
		}
		actions = append(actions, transferAction)

		engine := "mysql"
		if strings.Contains(containerName, "postgres") {
			engine = "postgres"
		}

		importArgs := []string{"exec", machine, "--", "bash", "/opt/nitro/scripts/docker-exec-import.sh", containerName, databaseName, filename, engine}
		dockerExecAction := nitro.Action{
			Type:       "exec",
			UseSyscall: false,
			Args:       importArgs,
		}
		actions = append(actions, dockerExecAction)

		fmt.Printf("Importing %q into %q (large files may take a while)...\n", filename, containerName)

		if err := nitro.Run(nitro.NewMultipassRunner("multipass"), actions); err != nil {
			return err
		}

		fmt.Println("Successfully imported the database backup into", containerName)

		return nil
	},
}
