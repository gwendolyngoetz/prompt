package prompt

import (
	"gwendolyngoetz/prompt/internal/icons"
	"gwendolyngoetz/prompt/pkg/computer"
	"gwendolyngoetz/prompt/pkg/git"
	p "gwendolyngoetz/prompt/pkg/prompt"
)

type Config struct {
	ShowHostname       bool
	ShowPythonVEnvName bool
}

var (
	colorBlack    = "0,0,0"
	colorWhite    = "224,224,224"
	colorOs       = "149,154,85"
	colorUserName = "225,85,85"
	colorHostname = "33,170,18"
	colorPwd      = "90,85,154"
	colorBranch   = "98,114,164"
	colorPython   = "55,118,171"
)

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
	return prompt.AddPart(p.Part{FgColor: colorBlack, BgColor: colorOs, Icon: computer.GetOsIcon(), Separator: icons.Separator})
}

func setSudo(prompt *p.PromptBuilder) {
	if computer.IsSudo() {
		prompt.AddPart(p.Part{FgColor: colorBlack, BgColor: colorUserName, Icon: icons.Sudo, Separator: icons.Separator})
	}
}

func setRemote(prompt *p.PromptBuilder, config *Config) {
	if computer.IsRemote() {
		hostnameText := ""

		if config.ShowHostname {
			hostnameText = computer.GetHostname()
		}

		prompt.AddPart(p.Part{FgColor: colorBlack, BgColor: colorHostname, Text: hostnameText, Icon: icons.Remote, Separator: icons.Separator})
	}
}

func setGit(prompt *p.PromptBuilder) {
	prompt.
		AddPart(p.Part{FgColor: colorBlack, BgColor: colorPwd, Text: git.GetRepoName(), Icon: icons.Git}).
		AddPart(p.Part{FgColor: colorWhite, BgColor: colorPwd, Text: git.GetRepoRelativePath(), Separator: icons.Separator}).
		AddPart(p.Part{FgColor: colorBlack, BgColor: colorBranch, Text: git.GetBranchName(), Icon: icons.Branch}).
		AddPart(p.Part{FgColor: colorBlack, BgColor: colorBranch, Icon: getChangesIcon(), Separator: icons.Separator})
}

func setPath(prompt *p.PromptBuilder) {
	prompt.AddPart(p.Part{FgColor: colorWhite, BgColor: colorPwd, Text: computer.GetPwd(), Icon: icons.Directory, Separator: icons.Separator})
}

func setPythonEnv(prompt *p.PromptBuilder, config *Config) {
	if computer.IsPythonVirtualEnv() {
		venvName := ""

		if config.ShowPythonVEnvName {
			venvName = computer.GetPythonVirtualEnv()
		}

		prompt.AddPart(p.Part{FgColor: colorBlack, BgColor: colorPython, Text: venvName, Icon: icons.Python, Separator: icons.Separator})
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
