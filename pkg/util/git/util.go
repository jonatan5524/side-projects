package util

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	core "github.com/jonatan5524/side-projects-manager/pkg/core/errors"
	util "github.com/jonatan5524/side-projects-manager/pkg/util/io"
)

const VERSION_CONTROL_FOLDER_NAME string = ".git"

func IsInVersionControl(path string) (bool, error) {
	folders, err := util.ListDirectory(path, util.FilterByDirectories)

	if err != nil {
		return false, core.NewIOError(path, err)
	}

	for _, file := range folders {
		if file.Name() == VERSION_CONTROL_FOLDER_NAME {
			return true, nil
		}
	}

	return false, nil
}

func executeCommand(app string, args ...string) (string, error) {
	cmd := exec.Command(app, args...)

	stdout, err := cmd.Output()

	if err != nil {
		return "-", core.NewExecError(fmt.Sprintf("%s args: %v", app, args), err)
	}

	return string(stdout), nil
}

func executeGitCommand(path string, args ...string) (string, error) {
	args = append([]string{"-C", path}, args...)

	return executeCommand("git", args...)
}

func VersionControlHostType(path string) (string, error) {
	args := []string{"config", "--get", "remote.origin.url"}

	url, err := executeGitCommand(path, args...)

	return ParseRemoteURL(url), err
}

func ParseRemoteURL(url string) string {
	// git@host:repository
	const SSH_CONNECTION_PREFIX = "git@"
	const SSH_HOST_REGEX_PATTERN = "@(.*?):"

	// https://host/repository
	const HTTPS_CONNECTION_PREFIX = "https://"
	const HTTPS_HOST_REGEX_PATTERN = "https:\\/\\/(.*?)\\/"

	var match []string

	if strings.HasPrefix(url, SSH_CONNECTION_PREFIX) {
		re := regexp.MustCompile(SSH_HOST_REGEX_PATTERN)
		match = re.FindStringSubmatch(url)
	} else if strings.HasPrefix(url, HTTPS_CONNECTION_PREFIX) {
		re := regexp.MustCompile(HTTPS_HOST_REGEX_PATTERN)
		match = re.FindStringSubmatch(url)
	}

	if len(match) < 2 {
		return "-"
	}

	return match[1]
}

func VersionControlCurrentBranch(path string) (string, error) {
	args := []string{"rev-parse", "--abbrev-ref", "HEAD"}

	return executeGitCommand(path, args...)
}

func IsVersionControlWorkingTreeClean(path string) (bool, error) {
	args := []string{"diff", "--exit-code"}

	answer, err := executeGitCommand(path, args...)

	isClean := answer == "" && err == nil

	return isClean, err
}
