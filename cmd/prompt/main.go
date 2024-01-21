package prompt

import (
	"gwendolyngoetz/prompt/internal/colors"
	"gwendolyngoetz/prompt/internal/icons"
	"gwendolyngoetz/prompt/pkg/computer"
	"gwendolyngoetz/prompt/pkg/git"
	p "gwendolyngoetz/prompt/pkg/prompt"
)

type Config struct {
	ShowHostname       bool
	ShowPythonVEnvName bool
}

var Version = "development"

func buildPrompt(config *Config) string {
	builder := p.PromptBuilder{}

	prompt := setOs(&builder)

	setSudo(prompt)
	setRemote(prompt, config)
	setPythonEnv(prompt, config)

	if git.IsRepo() {
		setGit(prompt)
	} else {
		setPath(prompt)
	}

	setPromptIcon(prompt)

	return prompt.Build()
}

func setOs(prompt *p.PromptBuilder) *p.PromptBuilder {
	return prompt.AddPart(p.Part{FgColor: colors.Black, BgColor: colors.Os, Icon: computer.GetOsIcon(), Separator: icons.Separator})
}

func setSudo(prompt *p.PromptBuilder) {
	if computer.IsSudo() {
		prompt.AddPart(p.Part{FgColor: colors.Black, BgColor: colors.UserName, Icon: icons.Sudo, Separator: icons.Separator})
	}
}

func setRemote(prompt *p.PromptBuilder, config *Config) {
	if computer.IsRemote() {
		hostnameText := ""

		if config.ShowHostname {
			hostnameText = computer.GetHostname()
		}

		prompt.AddPart(p.Part{FgColor: colors.Black, BgColor: colors.Hostname, Text: hostnameText, Icon: icons.Remote, Separator: icons.Separator})
	}
}

func setGit(prompt *p.PromptBuilder) {
	prompt.
		AddPart(p.Part{FgColor: colors.Black, BgColor: colors.Pwd, Text: git.GetRepoName(), Icon: icons.Git}).
		AddPart(p.Part{FgColor: colors.White, BgColor: colors.Pwd, Text: git.GetRepoRelativePath(), Separator: icons.Separator}).
		AddPart(p.Part{FgColor: colors.Black, BgColor: colors.Branch, Text: git.GetBranchName(), Icon: icons.Branch}).
		AddPart(p.Part{FgColor: colors.Black, BgColor: colors.Branch, Icon: getChangesIcon(), Separator: icons.Separator})
}

func setPath(prompt *p.PromptBuilder) {
	prompt.AddPart(p.Part{FgColor: colors.White, BgColor: colors.Pwd, Text: computer.GetPwd(), Icon: icons.Directory, Separator: icons.Separator})
}

func setPythonEnv(prompt *p.PromptBuilder, config *Config) {
	if computer.IsPythonVirtualEnv() {
		venvName := ""

		if config.ShowPythonVEnvName {
			venvName = computer.GetPythonVirtualEnv()
		}

		prompt.AddPart(p.Part{FgColor: colors.Black, BgColor: colors.Python, Text: venvName, Icon: icons.Python, Separator: icons.Separator})
	}
}

func setPromptIcon(prompt *p.PromptBuilder) {
	prompt.AddPart(p.Part{Icon: "\n" + icons.Prompt})
}

func getChangesIcon() string {
	if git.HasStatusChanges() {
		return icons.GitChanges
	}

	return ""
}
