package main

import (
	"flag"
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
	iconDirectory  = ""
	iconPrompt     = ""
)

var Version = "development"

func buildPrompt(showHostname bool) string {
	builder := p.PromptBuilder{}

	prompt := builder.AddPart(p.Part{FgColor: colorBlack, BgColor: colorOs, Icon: computer.GetOsIcon(), Separator: iconSeparator})

	if computer.IsSudo() {
		prompt.AddPart(p.Part{FgColor: colorBlack, BgColor: colorUserName, Icon: iconSudo, Separator: iconSeparator})
	}

	if computer.IsRemote() {
		hostnameText := ""

		if showHostname {
			hostnameText = computer.GetHostname()
		}

		prompt.AddPart(p.Part{FgColor: colorBlack, BgColor: colorHostname, Text: hostnameText, Icon: iconRemote, Separator: iconSeparator})
	}

	if git.IsRepo() {
		prompt.
			AddPart(p.Part{FgColor: colorBlack, BgColor: colorPwd, Text: git.GetRepoName(), Icon: iconGit}).
			AddPart(p.Part{FgColor: colorWhite, BgColor: colorPwd, Text: git.GetRepoRelativePath(), Separator: iconSeparator}).
			AddPart(p.Part{FgColor: colorBlack, BgColor: colorBranch, Text: git.GetBranchName(), Icon: iconBranch}).
			AddPart(p.Part{FgColor: colorBlack, BgColor: colorBranch, Icon: getChangesIcon(), Separator: iconSeparator})
	} else {
		prompt.AddPart(p.Part{FgColor: colorWhite, BgColor: colorPwd, Text: computer.GetPwd(), Icon: iconDirectory, Separator: iconSeparator})
	}

	prompt.AddPart(p.Part{Icon: "\n" + iconPrompt})

	return prompt.Build()
}

func getChangesIcon() string {
	if git.HasStatusChanges() {
		return iconGitChanges
	}

	return ""
}

func main() {
    showVersion := flag.Bool("version", false, "Show version")
    showHostname := flag.Bool("showHostname", false, "Show Hostname")
	flag.Parse()

	if *showVersion {
		fmt.Println(Version)
		return
	}

	promptString := buildPrompt(*showHostname)
	fmt.Println(promptString)
}
