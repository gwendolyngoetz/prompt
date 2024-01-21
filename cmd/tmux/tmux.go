package tmux

import (
	"fmt"
	"gwendolyngoetz/prompt/internal/icons"
	"gwendolyngoetz/prompt/pkg/computer"
	"gwendolyngoetz/prompt/pkg/git"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var format string

type TemplateData struct {
	OsIcon         string
	RepositoryName string
	BranchName     string
}

var Command = &cobra.Command{
	Use:   "tmux",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		formatValue := TemplateData{}

		if strings.Contains(format, ".OsIcon") {
			formatValue.OsIcon = computer.GetOsIcon()
		}

		if git.IsRepo() {
			if strings.Contains(format, ".RepositoryName") {
				formatValue.RepositoryName = fmt.Sprintf("%s %s ", icons.Git, git.GetRepoName())
			}

			if strings.Contains(format, ".BranchName") {
				formatValue.BranchName = fmt.Sprintf("%s %s ", icons.Branch, git.GetBranchName())
			}
		}

		tmpl, _ := template.New("format-test").Parse(format)
		_ = tmpl.Execute(os.Stdout, formatValue)
	},
}

func init() {
	Command.Flags().StringVarP(&format, "format", "", "{{.OsIcon}} {{.RepositoryName}} {{.BranchName}}", "")
}
