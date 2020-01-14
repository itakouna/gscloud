package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// makeConfigCmd represents the make-config command
var makeConfigCmd = &cobra.Command{
	Use:   "make-config",
	Short: "Create a new configuration file",
	Long: fmt.Sprintf(`Create a new and possibly almost empty configuration file overwriting an
existing one if it exists. Prints the path to the newly created file to
stdout.

Examples:

Create a new configuration file at the default configuration path
(%s/config.yaml):

	gscloud make-config

Create a new configuration file at a specified path:

	gscloud --config ~/myconfig.yaml make-config

`, cliConfigPath()),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := cliConfigPath()
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(path, emptyConfig(), 0644)
		if err != nil {
			return err
		}

		fmt.Println(path)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(makeConfigCmd)
}

// emptyConfig creates a new config YAML with a 'default account.
func emptyConfig() []byte {
	defaultAccount := accountEntry{
		Name:   "default",
		UserID: "",
		Token:  "",
		URL:    defaultAPIURL,
	}
	c := cliConfig{
		Accounts: []accountEntry{defaultAccount},
	}
	out, _ := yaml.Marshal(&c)
	return out
}
