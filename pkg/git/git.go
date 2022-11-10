package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func IsRepo() bool {
	_, err := exec.Command("git", "rev-parse", "--git-dir").Output()

	if werr, ok := err.(*exec.ExitError); ok {
		if s := werr.Error(); s != "0" {
			return false
		}
	}

	return true
}

func IsBareRepository() bool {
	out, _ := exec.Command(
		"git",
		"rev-parse",
		"--is-bare-repository").
		Output()

	result := strings.TrimSpace(string(out))
	resultB, _ := strconv.ParseBool(result)
	return resultB
}

func GetRepoName() string {
	out, _ := exec.Command(
		"git",
		"rev-parse",
		"--absolute-git-dir").
		Output()

	result := strings.TrimSpace(string(out))
	result = filepath.Dir(result)
	result = filepath.Base(result)

	return strings.TrimSpace(result)
}

func GetRepoRelativePath() string {
	out, _ := exec.Command(
		"git",
		"rev-parse",
		"--absolute-git-dir").
		Output()

	repoDir := strings.TrimSpace(string(out))
	repoDir = filepath.Dir(repoDir)

	dir, _ := os.Getwd()

	result := strings.Replace(dir, repoDir, "", -1)

	if len(result) > 0 {
		result = replaceAtIndex(result, ':', 0)
	}

	return result
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func HasStatusChanges() bool {
	out, _ := exec.Command(
		"git",
		"status",
		"--porcelain",
		"--branch",
		"--untracked-files=normal",
		"--ignore-submodules=all",
		"--show-stash",
		"--no-column").
		Output()

	outArray := strings.Split(string(out), "\n")
	result := len(outArray) - 2 // -1 for header, -1 for final newline

	return result > 0
}

func HasStashes() bool {
	out, _ := exec.Command(
		"git",
		"stash",
		"list").
		Output()

	outArray := strings.Split(string(out), "\n")
	result := len(outArray) - 1 // -1 for final newline

	return result > 0
}

func HasChanges() bool {
	return HasStatusChanges() || HasStashes()
}

func GetBranchName() string {
	branchName, _ := exec.Command(
		"git",
		"branch",
		"--show-current").
		Output()

	result := strings.TrimSpace(string(branchName))

	if len(result) > 0 {
		return result
	}

	commitHash, _ := exec.Command(
		"git",
		"rev-parse",
		"--short",
		"HEAD").
		Output()

	return strings.TrimSpace(string(commitHash))
}
