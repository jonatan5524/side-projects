package util

import (
	core "github.com/jonatan5524/side-projects-manager/pkg/core/errors"
	util "github.com/jonatan5524/side-projects-manager/pkg/util/io"
)

func IsInVersionControl(path string) (bool, error) {
	const VERSION_CONTROL_FOLDER_NAME string = ".git"
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
