package main

import (
	"flag"
	"fmt"
	"gwendolyngoetz/prompt/pkg/computer"
	"gwendolyngoetz/prompt/pkg/git"
	p "gwendolyngoetz/prompt/pkg/prompt"
)

type Config struct {
	ShowVersion        bool
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

var (
	iconSeparator  = ""
	iconSudo       = ""
	iconRemote     = ""
	iconGit        = ""
	iconBranch     = ""
	iconGitChanges = ""
	iconDirectory  = ""
	iconPrompt     = ""
	iconPython     = ""
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
	return prompt.AddPart(p.Part{FgColor: colorBlack, BgColor: colorOs, Icon: computer.GetOsIcon(), Separator: iconSeparator})
}

func setSudo(prompt *p.PromptBuilder) {
	if computer.IsSudo() {
		prompt.AddPart(p.Part{FgColor: colorBlack, BgColor: colorUserName, Icon: iconSudo, Separator: iconSeparator})
	}
}

func setRemote(prompt *p.PromptBuilder, config *Config) {
	if computer.IsRemote() {
		hostnameText := ""

		if config.ShowHostname {
			hostnameText = computer.GetHostname()
		}

		prompt.AddPart(p.Part{FgColor: colorBlack, BgColor: colorHostname, Text: hostnameText, Icon: iconRemote, Separator: iconSeparator})
	}
}

func setGit(prompt *p.PromptBuilder) {
	prompt.
		AddPart(p.Part{FgColor: colorBlack, BgColor: colorPwd, Text: git.GetRepoName(), Icon: iconGit}).
		AddPart(p.Part{FgColor: colorWhite, BgColor: colorPwd, Text: git.GetRepoRelativePath(), Separator: iconSeparator}).
		AddPart(p.Part{FgColor: colorBlack, BgColor: colorBranch, Text: git.GetBranchName(), Icon: iconBranch}).
		AddPart(p.Part{FgColor: colorBlack, BgColor: colorBranch, Icon: getChangesIcon(), Separator: iconSeparator})
}

func setPath(prompt *p.PromptBuilder) {
	prompt.AddPart(p.Part{FgColor: colorWhite, BgColor: colorPwd, Text: computer.GetPwd(), Icon: iconDirectory, Separator: iconSeparator})
}

func setPythonEnv(prompt *p.PromptBuilder, config *Config) {
	if computer.IsPythonVirtualEnv() {
		venvName := ""

		if config.ShowPythonVEnvName {
			venvName = computer.GetPythonVirtualEnv()
		}

		prompt.AddPart(p.Part{FgColor: colorBlack, BgColor: colorPython, Text: venvName, Icon: iconPython, Separator: iconSeparator})
	}
}

func setPromptIcon(prompt *p.PromptBuilder) {
	prompt.AddPart(p.Part{Icon: "\n" + iconPrompt})
}

func getChangesIcon() string {
	if git.HasStatusChanges() {
		return iconGitChanges
	}

	return ""
}

func loadConfig() *Config {
	config := Config{}
	flag.BoolVar(&config.ShowVersion, "version", false, "Show version")
	flag.BoolVar(&config.ShowHostname, "showHostname", false, "Show Hostname")
	flag.BoolVar(&config.ShowPythonVEnvName, "showPythonVEnvName", false, "Show Python Virtual Env Name")
	flag.Parse()

	return &config
}

func main() {
	config := loadConfig()

	if config.ShowVersion {
		fmt.Println(Version)
		return
	}

	promptString := buildPrompt(config)
	fmt.Println(promptString)
}
