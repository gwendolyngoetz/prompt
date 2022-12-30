package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type CloneType int

const (
	NoWorkTree CloneType = iota
	NonBareWorkTree
	BareWorkTree
)

func IsRepo() bool {
	_, err := exec.Command("git", "rev-parse", "--git-dir").Output()

	if werr, ok := err.(*exec.ExitError); ok {
		if s := werr.Error(); s != "0" {
			return false
		}
	}

	isBareRepo := IsBareRepository()

	if isBareRepo {
		return false
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

func IsInWorkTree() bool {
	out, _ := exec.Command(
		"git",
		"rev-parse",
		"--is-inside-work-tree").
		Output()

	result := strings.TrimSpace(string(out))
	resultB, _ := strconv.ParseBool(result)
	return resultB
}

func GetRepoName() string {
	cloneType := GetCloneType()

	if cloneType == NonBareWorkTree {
		return GetAbsoluteDir()
	}

	if cloneType == BareWorkTree {
		return GetCommonDir()
	}

	return ""
}

func GetAbsoluteDir() string {
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

func GetCommonDir() string {
	out, _ := exec.Command(
		"git",
		"rev-parse",
		"--git-common-dir").
		Output()

	result := strings.TrimSpace(string(out))
	// result = filepath.Dir(result)
	result = filepath.Base(result)

	return strings.TrimSpace(result)
}

func getTopLevel() string {
	out, _ := exec.Command(
		"git",
		"rev-parse",
		"--show-toplevel").
		Output()

	return strings.TrimSpace(string(out))
}

func GetCloneType() CloneType {
	topLevel := getTopLevel()

	_, err := os.Stat(topLevel)
	if err != nil {
		return NoWorkTree
	}

	info, err := os.Stat(topLevel + "/.git")
	if err != nil {
		return NoWorkTree
	}

	if info.IsDir() {
		return NonBareWorkTree
	}

	return BareWorkTree
}

func GetRepoRelativePath() string {
	out, _ := exec.Command(
		"git",
		"rev-parse",
		"--show-prefix").
		Output()

	result := strings.TrimSpace(string(out))
	result = strings.TrimSuffix(result, "/")

	if len(result) > 0 {
		result = ":" + result
	}

	return result
}

func GetRepoRelativePath1() string {
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
