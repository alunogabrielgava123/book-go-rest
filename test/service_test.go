package testcontroller

import (
	"fmt"
	"testing"
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/infra"
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/repository"
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/usecase"
)

func TestAlbumService(t *testing.T) {

	// Set default config
	dbpg := infra.ConfigPgSql()
	repo := repository.ImpReposityPgContructor(dbpg)


	t.Run("Delete album with id", func(t *testing.T) {

		//delete albums and
		expectValeu := 1

		service := usecase.DeleteAlbumByIdUseCaseConstructor(repo)
		value, err := service.Execute(21)
		if err != nil {
			t.Errorf("Nao foi possivel deletar esse id pois ele n existe")
		}

		if expectValeu == int(value) {
			fmt.Println("Deletando com sucesso!")
		}

	})

	// t.Run("Find all albums", func(t *testing.T) {

	// 	expectedSize := 1

	// 	service := usecase.FindeAlbunsUseCaseContrusctor(repo)

	// 	albums, err := service.Execute()
	// 	if err != nil {
	// 		t.Errorf("Nao foi possivel buscar todos so albums")
	// 	}

	// 	if len(albums) != expectedSize {
	// 		t.Errorf("Tamnho do album diferente do esperado!")
	// 	}
	// })

}
