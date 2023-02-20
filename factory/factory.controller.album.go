package factory

import (
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/controller"
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/infra"
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/repository"
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/usecase"
)

func FacotoryControllerAlbum() *controller.ControllerAlbum {

	dbpg := infra.ConfigPgSql()
	repo := repository.ImpReposityPgContructor(dbpg)

	//inject dependenci
	findAllUseCase := usecase.FindeAlbunsUseCaseContrusctor(repo)
	addAlbumUseCase := usecase.AddAlbunUseCaseConstructor(repo)
	deleteAlbumUseCase := usecase.DeleteAlbumByIdUseCaseConstructor(repo)

	return controller.ControllerAlbumConstructor(*addAlbumUseCase, *findAllUseCase, *deleteAlbumUseCase)
}
