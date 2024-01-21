package tmux

import (
	"fmt"
	"gwendolyngoetz/prompt/internal/icons"
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
		isRepo := git.IsRepo()

		if showRepoName == true && isRepo {
			message += fmt.Sprintf("%s %s ", icons.Git, git.GetRepoName())
		}

		if showGitBranch == true && isRepo {
			message += fmt.Sprintf("%s %s ", icons.Branch, git.GetBranchName())
		}

		fmt.Printf(strings.TrimSpace(message))
	},
}

func init() {
	Command.Flags().BoolVarP(&showGitBranch, "showGitBranch", "", false, "")
	Command.Flags().BoolVarP(&showRepoName, "showRepoName", "", false, "")
}
