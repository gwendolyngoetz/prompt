package tmux

import (
	"fmt"
	"gwendolyngoetz/prompt/pkg/git"
	"strings"

	"github.com/spf13/cobra"
)

var (
	showGitBranch bool
	showRepoName  bool
)

var Command = &cobra.Command{
	Use:   "tmux",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		message := ""

		if showRepoName == true {
			message += fmt.Sprintf("%s ", git.GetRepoName())
		}

		if showGitBranch == true {
			message += fmt.Sprintf("%s ", git.GetBranchName())
		}

		fmt.Printf(strings.TrimSpace(message))
	},
}

func init() {
	Command.Flags().BoolVarP(&showGitBranch, "showGitBranch", "", false, "")
	Command.Flags().BoolVarP(&showRepoName, "showRepoName", "", false, "")
}
