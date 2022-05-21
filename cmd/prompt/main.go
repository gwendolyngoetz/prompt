package main

import (
	"fmt"
	"gwendolyngoetz/prompt/pkg/computer"
	"gwendolyngoetz/prompt/pkg/git"
	p "gwendolyngoetz/prompt/pkg/prompt"
)

var (
	colorBlack    = "0,0,0"
	colorWhite    = "224,224,224"
	colorOs       = "149,154,85"
	colorUserName = "225,85,85"
	colorHostname = "33,170,18"
	colorPwd      = "90,85,154"
	colorBranch   = "98,114,164"
)

var (
	iconSeparator  = ""
	iconSudo       = ""
	iconRemote     = ""
	iconGit        = ""
	iconBranch     = ""
	iconGitChanges = ""
	iconDirectory  = "ﱮ"
	iconPrompt     = ""
)

func buildPrompt() string {
	builder := p.PromptBuilder{}

	prompt := builder.AddPart(p.Part{FgColor: colorBlack, BgColor: colorOs, Icon: computer.GetOsIcon(), Separator: iconSeparator})

	if computer.IsSudo() {
		prompt.AddPart(p.Part{FgColor: colorBlack, BgColor: colorUserName, Icon: iconSudo, Separator: iconSeparator})
	}

	if computer.IsRemote() {
		prompt.AddPart(p.Part{FgColor: colorBlack, BgColor: colorHostname, Icon: iconRemote, Separator: iconSeparator})
	}

	if git.IsRepo() {
		prompt.
			AddPart(p.Part{FgColor: colorBlack, BgColor: colorPwd, Text: git.GetRepoName(), Icon: iconGit}).
			AddPart(p.Part{FgColor: colorWhite, BgColor: colorPwd, Text: git.GetRepoRelativePath(), Separator: iconSeparator}).
			AddPart(p.Part{FgColor: colorBlack, BgColor: colorBranch, Text: git.GetBranchName(), Icon: iconBranch})

		gitChangesIcon := ""
		if git.HasChanges() {
			gitChangesIcon = iconGitChanges
		}

		prompt.AddPart(p.Part{FgColor: colorBlack, BgColor: colorBranch, Icon: gitChangesIcon, Separator: iconSeparator})
	} else {

		currentDir := computer.GetPwd()
		prompt.AddPart(p.Part{FgColor: colorWhite, BgColor: colorPwd, Text: currentDir, Icon: iconDirectory, Separator: iconSeparator})
	}

	prompt.AddPart(p.Part{Icon: "\n" + iconPrompt})

	return prompt.Build()
}

func main() {
	promptString := buildPrompt()
	fmt.Println(promptString)
}
