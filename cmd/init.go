package cmd

import (
	config "github.com/jonatan5524/side-projects-manager/pkg/config/db"
	DirectoryRepository "github.com/jonatan5524/side-projects-manager/pkg/repository/parentDirectory"
	ProjectRepository "github.com/jonatan5524/side-projects-manager/pkg/repository/project"
	DirectoryUsecase "github.com/jonatan5524/side-projects-manager/pkg/usecase/parentDirectory"
	ProjectUsecase "github.com/jonatan5524/side-projects-manager/pkg/usecase/project"
	"github.com/objectbox/objectbox-go/objectbox"
)

func initDB() *objectbox.ObjectBox {
	db, err := config.InitDB()

	if err != nil {
		panic(err)
	}

	return db
}

func initProjectUsecase(db *objectbox.ObjectBox) ProjectUsecase.ProjectUsecase {
	repository := ProjectRepository.NewProjectObjectBoxRepository(db)

	return ProjectUsecase.NewProjectService(repository)
}

func initParentDirectoryUsecase(db *objectbox.ObjectBox) DirectoryUsecase.ParentDirectoryUsecase {
	repository := DirectoryRepository.NewParentDirectoryObjectBoxRepository(db)

	return DirectoryUsecase.NewParentDirectoryService(repository)
}
