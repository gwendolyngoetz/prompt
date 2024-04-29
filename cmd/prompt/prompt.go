package prompt

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	toggle string
	config Config = Config{}
)

var Command = &cobra.Command{
	Use:   "prompt",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		promptString := buildPrompt(&config)
		fmt.Println(promptString)
	},
}

func init() {
	Command.Flags().BoolVarP(&config.HideOsIcon, "hideOsIcon", "", false, "Hide OS Icon")
	Command.Flags().BoolVarP(&config.ShowHostname, "showHostname", "", false, "Show Hostname")
	Command.Flags().BoolVarP(&config.ShowPythonVEnvName, "showPythonVEnvName", "", false, "Show Python Virtual Env Name")
	Command.Flags().BoolVarP(&config.HideGit, "hideGit", "", false, "Show Git Repo and Branch Names")
}
